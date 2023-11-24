import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import {
  HiddenResource,
  InventoryFilterData,
  InventoryItem,
  InventoryStats,
  View
} from '../types/useInventoryTypes';
import parseURLParams from './parseURLParams';

type getCustomViewInventoryListAndStatsProps = {
  router: NextRouter;
  views: View[] | undefined;
  setSearchedLoading: (searchedLoading: boolean) => void;
  setStatsLoading: (statsLoading: boolean) => void;
  showToast: (value: ToastProps) => void;
  setError: (error: boolean) => void;
  setHiddenResources: (
    value: SetStateAction<HiddenResource[] | undefined>
  ) => void;
  setInventoryStats: (inventoryStats: InventoryStats | undefined) => void;
  batchSize: number;
  setSearchedInventory: (
    searchedInventory: InventoryItem[] | undefined
  ) => void;
  setDisplayedFilters: (
    displayedFilters: InventoryFilterData[] | undefined
  ) => void;
  setSkippedSearch: (value: SetStateAction<number>) => void;
  setHideOrUnhideHasUpdate: (value: SetStateAction<boolean>) => void;
  setShouldFetchMore: (value: SetStateAction<boolean>) => void;
};

/** Fetch inventory, top stats and hidden resources for a given custom view.
 * - Inventory list will be stored in the state: searchedInventory
 * - Inventory top stats will be stored in the state: inventoryStats
 */
function getCustomViewInventoryListAndStats({
  router,
  views,
  setSearchedLoading,
  setStatsLoading,
  showToast,
  setError,
  setHiddenResources,
  setInventoryStats,
  batchSize,
  setSearchedInventory,
  setDisplayedFilters,
  setSkippedSearch,
  setHideOrUnhideHasUpdate,
  setShouldFetchMore
}: getCustomViewInventoryListAndStatsProps) {
  if (router.query.view && views && views.length > 0) {
    const id = router.query.view;
    const filterFound = views.find(view => view.id.toString() === id);

    if (filterFound) {
      setSearchedLoading(true);
      setStatsLoading(true);
      const payloadJson = JSON.stringify(filterFound?.filters);

      settingsService.getHiddenResourcesFromView(id as string).then(res => {
        if (res === Error) {
          setError(true);
        } else {
          setHiddenResources(res);
        }
      });

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
        .getInventory(`?limit=${batchSize}&skip=0&view=${id}`, payloadJson)
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
            setSkippedSearch(prev => prev + batchSize);
            setSearchedLoading(false);
            const newFiltersToDisplay: InventoryFilterData[] =
              filterFound.filters.map(filter =>
                parseURLParams(filter, 'display', true)
              );
            setDisplayedFilters(newFiltersToDisplay);
            setHideOrUnhideHasUpdate(false);

            if (res.length >= batchSize) {
              setShouldFetchMore(true);
            } else {
              setShouldFetchMore(false);
            }
          }
        });
    } 
    // TODO: https://github.com/tailwarden/komiser/issues/1208
    // else {
    //   setTimeout(() => router.push(router.pathname), 5000);
    //   return showToast({
    //     hasError: true,
    //     title: `Invalid view`,
    //     message: `We couldn't find this view. Redirecting back to home...`
    //   });
    // }
  }
  return null;
}

export default getCustomViewInventoryListAndStats;
