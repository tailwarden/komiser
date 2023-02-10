import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import {
  InventoryFilterData,
  InventoryItem,
  InventoryStats
} from '../types/useInventoryTypes';

type getInventoryListAndStatsProps = {
  router: NextRouter;
  setStatsLoading: (statsLoading: boolean) => void;
  setDisplayedFilters: (
    displayedFilters: InventoryFilterData[] | undefined
  ) => void;
  setFilters: (filters: InventoryFilterData[] | undefined) => void;
  setError: (error: boolean) => void;
  setSearchedInventory: (
    searchedInventory: InventoryItem[] | undefined
  ) => void;
  setInventoryStats: (inventoryStats: InventoryStats | undefined) => void;
  batchSize: number;
  setInventory: (inventory: InventoryItem[] | undefined) => void;
  setSkipped: (value: SetStateAction<number>) => void;
};

/** Fetch base inventory and top stats for All Resources.
 * - Inventory list will be stored in the state: inventory
 * - Inventory top stats will be stored in the state: inventoryStats
 */
function getInventoryListAndStats({
  router,
  setStatsLoading,
  setDisplayedFilters,
  setFilters,
  setError,
  setSearchedInventory,
  setInventoryStats,
  batchSize,
  setInventory,
  setSkipped
}: getInventoryListAndStatsProps) {
  if (
    router.query &&
    Object.keys(router.query).length === 0 &&
    !router.query.view
  ) {
    setStatsLoading(true);
    setDisplayedFilters(undefined);
    setSearchedInventory(undefined);
    setFilters(undefined);

    settingsService.getInventoryStats().then(res => {
      if (res === Error) {
        setError(true);
        setStatsLoading(false);
      } else {
        setInventoryStats(res);
        setStatsLoading(false);
      }
    });

    settingsService.getInventory(`?limit=${batchSize}&skip=0`).then(res => {
      if (res === Error) {
        setError(true);
      } else {
        setInventory(res);
        setSkipped(prev => prev + batchSize);
      }
    });
  }
}

export default getInventoryListAndStats;
