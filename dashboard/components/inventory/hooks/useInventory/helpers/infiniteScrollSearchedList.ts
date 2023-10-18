import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import { InventoryItem } from '../types/useInventoryTypes';

type InfiniteScrollSearchedListProps = {
  shouldFetchMore: boolean;
  isVisible: boolean;
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
function infiniteScrollSearchedList({
  shouldFetchMore,
  isVisible,
  query,
  router,
  batchSize,
  skippedSearch,
  showToast,
  setSearchedInventory,
  setShouldFetchMore,
  setSkippedSearch
}: InfiniteScrollSearchedListProps) {
  if (
    shouldFetchMore &&
    isVisible &&
    query &&
    Object.keys(router.query).length === 0
  ) {
    settingsService
      .getInventory(`?limit=${batchSize}&skip=${skippedSearch}&query=${query}`)
      .then(res => {
        if (res === Error) {
          showToast({
            hasError: true,
            title: `There was an error fetching more resources!`,
            message: `Please refresh the page and try again.`
          });
        }

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
      });
  }
}

export default infiniteScrollSearchedList;
