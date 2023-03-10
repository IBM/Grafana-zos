import produce from 'immer';
import { ResourceRow, ResourceRowGroup } from './types';

// This regex matches URIs representing:
//  - subscriptions: /subscriptions/44693801-6ee6-49de-9b2d-9106972f9572
//  - resource groups: /subscriptions/44693801-6ee6-49de-9b2d-9106972f9572/resourceGroups/cloud-datasources
//  - resources: /subscriptions/44693801-6ee6-49de-9b2d-9106972f9572/resourceGroups/cloud-datasources/providers/Microsoft.Compute/virtualMachines/GithubTestDataVM
const RESOURCE_URI_REGEX =
  /\/subscriptions\/(?<subscriptionID>[^/]+)(?:\/resourceGroups\/(?<resourceGroup>[^/]+)(?:\/providers.+\/(?<resource>[^/]+))?)?/;

type RegexGroups = Record<string, string | undefined>;

export function parseResourceURI(resourceURI: string) {
  const matches = RESOURCE_URI_REGEX.exec(resourceURI);
  const groups: RegexGroups = matches?.groups ?? {};
  const { subscriptionID, resourceGroup, resource } = groups;

  if (!subscriptionID) {
    return undefined;
  }

  return { subscriptionID, resourceGroup, resource };
}

export function isGUIDish(input: string) {
  return !!input.match(/^[A-Z0-9]+/i);
}

export function findRow(rows: ResourceRowGroup, uri: string): ResourceRow | undefined {
  for (const row of rows) {
    if (row.uri.toLowerCase() === uri.toLowerCase()) {
      return row;
    }

    if (row.children) {
      const result = findRow(row.children, uri);

      if (result) {
        return result;
      }
    }
  }

  return undefined;
}

export function addResources(rows: ResourceRowGroup, targetParentId: string, newResources: ResourceRowGroup) {
  return produce(rows, (draftState) => {
    const draftRow = findRow(draftState, targetParentId);

    // we can't find the selected resource in our list of resources,
    // probably means user has either mistyped in the input field
    // or is using template variables.
    // either way no need to throw, just show that none of the resources are checked
    if (!draftRow) {
      return;
    }

    draftRow.children = newResources;
  });
}
