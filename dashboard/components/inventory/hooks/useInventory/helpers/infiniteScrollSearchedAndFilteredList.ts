import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import { InventoryFilterData, InventoryItem } from '../types/useInventoryTypes';

type InfiniteScrollSearchedAndFilteredList = {
  shouldFetchMore: boolean;
  isVisible: boolean;
  filters: InventoryFilterData[] | undefined;
  query: string;
  router: NextRouter;
  batchSize: number;
  skippedSearch: number;
  showToast: (value: ToastProps) => void;
  setSearchedInventory: (
    value: SetStateAction<InventoryItem[] | undefined>
  ) => void;
  setShouldFetchMore: (value: SetStateAction<boolean>) => void;
  setSkippedSearch: (value: SetStateAction<number>) => void;
};

/** Load the next 50 results when the user scrolls a searched inventory list to the end */
function infiniteScrollSearchedAndFilteredList({
  shouldFetchMore,
  isVisible,
  filters,
  query,
  router,
  batchSize,
  skippedSearch,
  showToast,
  setSearchedInventory,
  setShouldFetchMore,
  setSkippedSearch
}: InfiniteScrollSearchedAndFilteredList) {
  if (
    shouldFetchMore &&
    isVisible &&
    query &&
    Object.keys(router.query).length > 0 &&
    !router.query.view
  ) {
    let payloadJson = '';

    if (!router.query.view && filters && filters.length > 0) {
      payloadJson = JSON.stringify(filters);
    }

    settingsService
      .getInventory(
        `?limit=${batchSize}&skip=${skippedSearch}&query=${query}`,
        payloadJson
      )
      .then(res => {
        if (res.error) {
          showToast({
            hasError: true,
            title: `There was an error fetching more resources!`,
            message: `Please refresh the page and try again.`
          });
        } else {
          setSearchedInventory(prev => {
            if (prev) {
              return [...prev, ...res];
            }
            return res;
          });

          if (res.length >= batchSize) {
            setShouldFetchMore(true);
          } else {
            setShouldFetchMore(false);
          }

          setSkippedSearch(prev => prev + batchSize);
        }
      });
  }
}

export default infiniteScrollSearchedAndFilteredList;
