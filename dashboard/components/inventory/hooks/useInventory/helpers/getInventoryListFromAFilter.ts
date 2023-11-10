import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import {
  InventoryFilterData,
  InventoryItem,
  InventoryStats
} from '../types/useInventoryTypes';
import parseURLParams from './parseURLParams';

type getInventoryListFromAFilterProps = {
  router: NextRouter;
  setSearchedLoading: (searchedLoading: boolean) => void;
  setStatsLoading: (statsLoading: boolean) => void;
  showToast: (value: ToastProps) => void;
  setError: (error: boolean) => void;
  setInventoryStats: (inventoryStats: InventoryStats | undefined) => void;
  batchSize: number;
  setSearchedInventory: (
    searchedInventory: InventoryItem[] | undefined
  ) => void;
  setFilters: (filters: InventoryFilterData[] | undefined) => void;
  setDisplayedFilters: (
    displayedFilters: InventoryFilterData[] | undefined
  ) => void;
  setSkippedSearch: (value: SetStateAction<number>) => void;
  setShouldFetchMore: (value: SetStateAction<boolean>) => void;
};

/** Fetch inventory from a filter.
 * - Inventory list will be stored in the state: searchedInventory
 */
function getInventoryListFromAFilter({
  router,
  setSearchedLoading,
  setStatsLoading,
  showToast,
  setError,
  setInventoryStats,
  batchSize,
  setSearchedInventory,
  setFilters,
  setDisplayedFilters,
  setSkippedSearch,
  setShouldFetchMore
}: getInventoryListFromAFilterProps) {
  if (
    router.query &&
    Object.keys(router.query).length > 0 &&
    !router.query.view
  ) {
    if (
      Object.keys(router.query)[0].split(':').length <= 1 &&
      !router.query.view
    ) {
      setTimeout(() => router.push(router.pathname), 5000);
      return showToast({
        hasError: true,
        title: `Invalid URL params`,
        message: `There was an error processing the page. Redirecting back to home...`
      });
    }
    setSearchedLoading(true);
    setStatsLoading(true);

    const newFilters: InventoryFilterData[] = Object.keys(router.query).map(
      param => parseURLParams(param as string, 'fetch')
    );
    const newFiltersToDisplay: InventoryFilterData[] = Object.keys(
      router.query
    ).map(param => parseURLParams(param as string, 'display'));

    const payloadJson = JSON.stringify(newFilters);

    settingsService.getFilteredInventoryStats(payloadJson).then(res => {
      if (res === Error) {
        setError(true);
        setStatsLoading(false);
      } else {
        setInventoryStats(res);
        setStatsLoading(false);
      }
    });

    settingsService
      .getInventory(`?limit=${batchSize}&skip=0`, payloadJson)
      .then(res => {
        if (res.error) {
          showToast({
            hasError: true,
            title: `Filter could not be applied!`,
            message: `Please refresh the page and try again.`
          });
          setError(true);
          setSearchedLoading(false);
        } else {
          setSearchedInventory(res);
          setFilters(newFilters);
          setDisplayedFilters(newFiltersToDisplay);
          setSkippedSearch(prev => prev + batchSize);
          setSearchedLoading(false);

          if (res.length >= batchSize) {
            setShouldFetchMore(true);
          } else {
            setShouldFetchMore(false);
          }
        }
      });
  }
  return null;
}

export default getInventoryListFromAFilter;
