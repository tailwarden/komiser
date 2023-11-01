import { Provider } from '../../../../../utils/providerHelper';

export type InventoryFilterData = {
  field:
    | 'provider'
    | 'region'
    | 'account'
    | 'name'
    | 'service'
    | 'cost'
    | 'tags'
    | 'tag'
    | string
    | undefined;
  operator:
    | 'IS'
    | 'IS_NOT'
    | 'CONTAINS'
    | 'NOT_CONTAINS'
    | 'IS_EMPTY'
    | 'IS_NOT_EMPTY'
    | 'EXISTS'
    | 'NOT_EXISTS'
    | string
    | undefined;
  tagKey?: string;
  values: [] | string[];
};

export type InventoryStats = {
  resources: number;
  costs: number;
  savings: number;
  regions: number;
};

export type Tag = {
  key: string;
  value: string;
};

export type InventoryItem = {
  relations: any[];
  account: string;
  accountId: string;
  cost: number;
  createdAt: string;
  fetchedAt: string;
  id: string;
  link: string;
  metadata: null;
  name: string;
  provider: Provider;
  region: string;
  resourceId: string;
  service: string;
  tags: Tag[] | [] | null;
};
export type Pages = 'resource details' | 'tags' | 'delete';

export type View = {
  id: number;
  name: string;
  filters: InventoryFilterData[];
  exclude: string[];
};

export type HiddenResource = {
  id: number;
  resourceId: string;
  provider: string;
  account: string;
  accountId: string;
  service: string;
  region: string;
  name: string;
  createdAt: string;
  fetchedAt: string;
  cost: number;
  metadata: null;
  tags: Tag[] | [] | null;
  link: string;
  Value: string;
};
