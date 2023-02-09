import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import {
  InventoryFilterData,
  InventoryItem,
  InventoryStats
} from '../types/useInventoryTypes';

type InfiniteScrollInventoryListProps = {
  inventory: InventoryItem[] | undefined;
  inventoryStats: InventoryStats | undefined;
  skipped: number;
  isVisible: boolean;
  query: string;
  filters: InventoryFilterData[] | undefined;
  router: NextRouter;
  setError: (error: boolean) => void;
  batchSize: number;
  setInventory: (value: SetStateAction<InventoryItem[] | undefined>) => void;
  setSkipped: (value: SetStateAction<number>) => void;
};

/** Load the next 50 results when the user scrolls the inventory list to the end */
function infiniteScrollInventoryList({
  inventory,
  inventoryStats,
  skipped,
  isVisible,
  query,
  filters,
  router,
  setError,
  batchSize,
  setInventory,
  setSkipped
}: InfiniteScrollInventoryListProps) {
  if (
    inventory &&
    inventory.length > 0 &&
    inventoryStats &&
    skipped < inventoryStats.resources &&
    isVisible &&
    !query &&
    !filters &&
    !router.query.view
  ) {
    setError(false);

    settingsService
      .getInventory(`?limit=${batchSize}&skip=${skipped}`)
      .then(res => {
        if (res === Error) {
          setError(true);
        } else {
          setInventory(prev => {
            if (prev) {
              return [...prev, ...res];
            }
            return res;
          });
          setSkipped(prev => prev + batchSize);
        }
      });
  }
}

export default infiniteScrollInventoryList;
