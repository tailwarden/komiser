import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import { InventoryItem, View } from '../types/useInventoryTypes';

type InfiniteScrollCustomViewListProps = {
  router: NextRouter;
  shouldFetchMore: boolean;
  isVisible: boolean;
  views: View[] | undefined;
  query: string;
  batchSize: number;
  skippedSearch: number;
  showToast: (value: ToastProps) => void;
  setSearchedInventory: (
    value: SetStateAction<InventoryItem[] | undefined>
  ) => void;
  setShouldFetchMore: (value: SetStateAction<boolean>) => void;
  setSkippedSearch: (value: SetStateAction<number>) => void;
};

/** Load the next 50 results when the user scrolls a custom view inventory list to the end */
function infiniteScrollCustomViewList({
  router,
  shouldFetchMore,
  isVisible,
  views,
  query,
  batchSize,
  skippedSearch,
  showToast,
  setSearchedInventory,
  setShouldFetchMore,
  setSkippedSearch
}: InfiniteScrollCustomViewListProps) {
  if (
    shouldFetchMore &&
    isVisible &&
    views &&
    views.length > 0 &&
    router.query.view &&
    !query
  ) {
    const id = router.query.view;
    const filterFound = views.find(view => view.id.toString() === id);

    if (filterFound) {
      const payloadJson = JSON.stringify(filterFound?.filters);

      settingsService
        .getInventory(
          `?limit=${batchSize}&skip=${skippedSearch}&view=${id}`,
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
            setSkippedSearch(prev => prev + batchSize);

            if (res.length >= batchSize) {
              setShouldFetchMore(true);
            } else {
              setShouldFetchMore(false);
            }
          }
        });
    }
  }
}

export default infiniteScrollCustomViewList;
