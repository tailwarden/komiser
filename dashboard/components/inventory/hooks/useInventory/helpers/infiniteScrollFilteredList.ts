import { ToastProps } from '@components/toast/Toast';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import { InventoryFilterData, InventoryItem } from '../types/useInventoryTypes';

type InfiniteScrollFilteredListProps = {
  shouldFetchMore: boolean;
  isVisible: boolean;
  filters: InventoryFilterData[] | undefined;
  query: string;
  batchSize: number;
  skippedSearch: number;
  showToast: (value: ToastProps) => void;
  setQuery: (value: SetStateAction<string>) => void;
  setSearchedInventory: (
    value: SetStateAction<InventoryItem[] | undefined>
  ) => void;
  setShouldFetchMore: (value: SetStateAction<boolean>) => void;
  setSkippedSearch: (value: SetStateAction<number>) => void;
};

/** Load the next 50 results when the user scrolls a filtered inventory list to the end */
function infiniteScrollFilteredList({
  shouldFetchMore,
  isVisible,
  filters,
  query,
  batchSize,
  skippedSearch,
  showToast,
  setQuery,
  setSearchedInventory,
  setShouldFetchMore,
  setSkippedSearch
}: InfiniteScrollFilteredListProps) {
  if (shouldFetchMore && isVisible && filters && !query) {
    const payloadJson = JSON.stringify(filters);
    settingsService
      .getInventory(`?limit=${batchSize}&skip=${skippedSearch}`, payloadJson)
      .then(res => {
        if (res.error) {
          showToast({
            hasError: true,
            title: `There was an error fetching more resources!`,
            message: `Please refresh the page and try again.`
          });
        } else {
          setQuery('');
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

export default infiniteScrollFilteredList;
