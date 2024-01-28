import { useToast } from '@components/toast/ToastProvider';
import { useRouter } from 'next/router';
import { ChangeEvent, useEffect, useRef, useState } from 'react';
import {
  checkIfServiceIsSupported,
  checkIsSomeServiceUnavailable
} from '@utils/serviceHelper';
import settingsService from '../../../../services/settingsService';
import useIsVisible from '../useIsVisible/useIsVisible';
import getCustomViewInventoryListAndStats from './helpers/getCustomViewInventoryListAndStats';
import getInventoryListAndStats from './helpers/getInventoryListAndStats';
import getInventoryListFromAFilter from './helpers/getInventoryListFromAFilter';
import infiniteScrollCustomViewList from './helpers/infiniteScrollCustomViewList';
import infiniteScrollFilteredList from './helpers/infiniteScrollFilteredList';
import infiniteScrollInventoryList from './helpers/infiniteScrollInventoryList';
import infiniteScrollSearchedAndFilteredList from './helpers/infiniteScrollSearchedAndFilteredList';
import infiniteScrollSearchedCustomViewList from './helpers/infiniteScrollSearchedCustomViewList';
import infiniteScrollSearchedList from './helpers/infiniteScrollSearchedList';
import {
  HiddenResource,
  InventoryFilterData,
  InventoryItem,
  InventoryStats,
  Pages,
  Tag,
  View
} from './types/useInventoryTypes';

function useInventory() {
  const [isSomeServiceUnavailable, setIsSomeServiceUnavailable] =
    useState<boolean>(false);
  const [inventoryStats, setInventoryStats] = useState<InventoryStats>();
  const [inventory, setInventory] = useState<InventoryItem[]>();
  const [error, setError] = useState(false);
  const [skipped, setSkipped] = useState(0);
  const [skippedSearch, setSkippedSearch] = useState(0);
  const [inventoryHasUpdate, setInventoryHasUpdate] = useState(false);
  const [query, setQuery] = useState('');
  const [searchedInventory, setSearchedInventory] = useState<InventoryItem[]>();
  const [shouldFetchMore, setShouldFetchMore] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const [data, setData] = useState<InventoryItem>();
  const [page, setPage] = useState<Pages>('resource details');
  const [tags, setTags] = useState<Tag[]>();
  const [loading, setLoading] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [bulkItems, setBulkItems] = useState<string[] | []>([]);
  const [bulkSelectCheckbox, setBulkSelectCheckbox] = useState(false);
  const [displayedFilters, setDisplayedFilters] =
    useState<InventoryFilterData[]>();
  const [filters, setFilters] = useState<InventoryFilterData[]>();
  const [searchedLoading, setSearchedLoading] = useState(false);
  const [statsLoading, setStatsLoading] = useState(false);
  const [views, setViews] = useState<View[]>();
  const [hiddenResources, setHiddenResources] = useState<HiddenResource[]>();
  const [hideResourcesLoading, setHideResourcesLoading] = useState(false);
  const [hideOrUnhideHasUpdate, setHideOrUnhideHasUpdate] = useState(false);

  const { toast, showToast, dismissToast } = useToast();
  const reloadDiv = useRef<HTMLDivElement>(null);
  const isVisible = useIsVisible(reloadDiv);
  const batchSize: number = 50;
  const router = useRouter();

  /*
    Check if there are items in searchedInventory:
    - If yes, check if even one item is not supported.
    - If unsupported item found, set isSomeServiceUnavailable to true.
    - If searchedInventory is empty, get all services and check if even one is not supported.
      - If unsupported service found, set isSomeServiceUnavailable to true.
  */
  useEffect(() => {
    if (searchedInventory) {
      setIsSomeServiceUnavailable(
        searchedInventory.some(
          item => !checkIfServiceIsSupported(item.provider, item.service)
        )
      );
    } else {
      settingsService.getServices().then(res => {
        setIsSomeServiceUnavailable(checkIsSomeServiceUnavailable(res));
      });
    }
  }, [searchedInventory]);

  /** Reset most of the UI states:
   * - skipped (used to skip results in the data fetch call)
   * - skippedSearch (same, but used to skip results in the searched data fetch call)
   * - bulkItems (array used to nest all resource item ids selected)
   * - bulkSelectionCheckbox (boolean that handles all checkboxes checked/unchecked toggle)
   * - query (search value)
   * - filters (array used to nest all filters)
   * - displayedFilters (same, but used to nest all filters that compose the UI)
   * - shouldFetchMore (boolean that indicates whether or not the infinite scroll fetch should be triggered)
   */
  function resetStates() {
    setSkipped(0);
    setSkippedSearch(0);
    setBulkItems([]);
    setBulkSelectCheckbox(false);
    setQuery('');
    setFilters(undefined);
    setDisplayedFilters(undefined);
    setShouldFetchMore(false);
  }

  /** Fetch all the custom views.
   * - Views will be stored in the state: views.
   */
  function getViews(edit?: boolean, viewId?: string, redirect?: boolean) {
    settingsService.getViews().then(res => {
      if (res === Error) {
        showToast({
          hasError: true,
          title: `The custom views couldn't be loaded.`,
          message: `There was a problem fetching the views. Please try again.`
        });
        setError(true);
      } else {
        const sortViewsById: View[] = res;
        sortViewsById.sort((a, b) => a.id - b.id);
        setViews(sortViewsById);

        if (edit && viewId) {
          router.push(`?view=${viewId}`, undefined, { shallow: true });
        }

        if (redirect) {
          const idFromNewCustomView = res[res.length - 1].id;
          router.push(`?view=${idFromNewCustomView}`, undefined, {
            shallow: true
          });
        }
      }
    });
  }

  /** Whenever the page changes or views is updated, fetch the custom views, reset UI states and fetch the inventory based on the URL params */
  useEffect(() => {
    resetStates();

    if (!views) {
      getViews();
    }

    if (views) {
      getInventoryListAndStats({
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
      });

      getInventoryListFromAFilter({
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
      });

      getCustomViewInventoryListAndStats({
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
      });
    }
  }, [router.query, views]);

  /** Infinite scrolling handler. Identifies which inventory should be fetched on scroll. */
  useEffect(() => {
    infiniteScrollInventoryList({
      inventory,
      inventoryStats,
      skipped,
      isVisible,
      query,
      filters,
      router,
      batchSize,
      showToast,
      setInventory,
      setSkipped
    });

    infiniteScrollFilteredList({
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
    });

    infiniteScrollSearchedList({
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
    });

    infiniteScrollSearchedAndFilteredList({
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
    });

    infiniteScrollCustomViewList({
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
    });

    infiniteScrollSearchedCustomViewList({
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
    });
  }, [isVisible]);

  /** Search effect behavior:
   * - If there's a filtered list, search should only bring back results from the list
   * - If not, search should get from all inventory
   * - A filter can overwrite a search, but not the opposite */
  useEffect(() => {
    let mounted = true;

    setSkippedSearch(0);
    setShouldFetchMore(false);
    setBulkItems([]);
    setBulkSelectCheckbox(false);

    let payload: any;
    let payloadJson: string;
    let id: string;

    // If this is the all resources and search value is empty, fetch again the inventory list
    if (views && !query && !router.query.view && !filters) {
      getInventoryListAndStats({
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
      });
    }

    // If this is the all resources filtered and search value is empty, fetch again the filtered inventory list
    if (
      views &&
      !query &&
      !router.query.view &&
      filters &&
      filters.length > 0
    ) {
      getInventoryListFromAFilter({
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
      });
    }

    // If this is a custom view and search value is empty, fetch again the custom view list
    if (views && !query && router.query.view) {
      getCustomViewInventoryListAndStats({
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
      });
    }

    // When query has a value, handle the payload based on the URL params
    if (query) {
      setSearchedLoading(true);

      // All resources search payload handler
      if (!filters && query && !router.query.view) {
        payload = [];
      }

      // All resources + filter search payload handler
      if (filters && filters.length > 0 && !router.query.view) {
        payload = filters;
      }

      // Custom view search payload handler
      if (router.query.view && views && views.length > 0) {
        id = router.query.view.toString();
        const filterFound = views.find(
          view => view.id.toString() === router.query.view
        );

        if (filterFound) {
          payload = filterFound.filters;
        }
      }

      setTimeout(() => {
        payloadJson = JSON.stringify(payload);
        if (mounted) {
          settingsService
            .getInventory(
              `?limit=${batchSize}&skip=0&query=${query}${
                id ? `&view=${id}` : ''
              }`,
              payloadJson
            )
            .then(res => {
              if (mounted) {
                if (res.error) {
                  showToast({
                    hasError: true,
                    title: `Search error`,
                    message: `There was an error searching for ${query}. Please refer to the logs and try again.`
                  });
                  setError(true);
                  setSearchedLoading(false);
                } else {
                  setSearchedInventory(res);
                  setSkippedSearch(prev => prev + batchSize);
                  setSearchedLoading(false);

                  if (res.length >= batchSize) {
                    setShouldFetchMore(true);
                  } else {
                    setShouldFetchMore(false);
                  }
                }
              }
            });
        }
      }, 700);
    }

    return () => {
      mounted = false;
    };
  }, [query]);

  /** When a resource is hid or unhid, reset the states and call getCustomViewInventoryListAndStats. */
  useEffect(() => {
    if (hideOrUnhideHasUpdate) {
      resetStates();
      getCustomViewInventoryListAndStats({
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
      });
    }
  }, [hideOrUnhideHasUpdate]);

  /** Refresh list when tags are saved.
   * - If it's on all resources, refetch the inventory list
   * - If it's on a custom view, quick reload the current url
   */
  useEffect(() => {
    let mounted = true;

    if (inventoryHasUpdate) {
      if (Object.keys(router.query).length === 0) {
        settingsService.getInventory(`?limit=${batchSize}&skip=0`).then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
              showToast({
                hasError: true,
                title: `There was an error re-fetching the resources!`,
                message: `Please refresh the page and try again.`
              });
            } else {
              setQuery('');
              setInventory(res);
              setSkipped(batchSize);
              setInventoryHasUpdate(false);
            }
          }
        });
      } else {
        setSkippedSearch(0);
        setInventoryHasUpdate(false);
        router.push(router.asPath, undefined, { shallow: true });
      }
    }

    return () => {
      mounted = false;
    };
  }, [inventoryHasUpdate]);

  /** Clean the modal:
   * - Sets the state data to undefined
   * - Sets the page to 'tags'
   */
  function cleanModal() {
    setData(undefined);
    setPage('resource details');
  }

  /** Opens the modal, as well as:
   * - Calls cleanModal first
   * - Sets the data to the inventory item passed as argument
   * - If the inventory item has tags, set the tags state to them.
   * - If the inventory item has no tags, set the tags state to an array with an object inside containing empty key and value
   */
  function openModal(inventoryItem: InventoryItem) {
    cleanModal();
    setData(inventoryItem);

    if (inventoryItem.tags && inventoryItem.tags.length > 0) {
      setTags(inventoryItem.tags);
    } else {
      setTags([{ key: '', value: '' }]);
    }

    setIsOpen(true);
  }

  /** Opens the modal for bulk operations, as well as:
   * - Calls cleanModal first
   * - Set the tags state to an array with an object inside containing empty key and value
   */
  function openBulkModal() {
    cleanModal();
    setTags([{ key: '', value: '' }]);
    setIsOpen(true);
  }

  /** Close the modal by setting the isOpen to false */
  function closeModal() {
    setIsOpen(false);
  }

  /** Handles the page change inside the modal */
  function goTo(newPage: Pages) {
    setPage(newPage);
  }

  /** Handles the change for the key and value inputs inside the modal */
  function handleChange(newData: Partial<Tag>, id?: number) {
    if (tags && typeof id === 'number') {
      const newValues: Tag[] = [...tags];
      newValues[id] = {
        ...newValues[id],
        ...newData
      };
      setTags(newValues);
    }
  }

  /** Creates another key and value object inside the tags array */
  function addNewTag() {
    if (tags) {
      setTags(prev => {
        if (prev) {
          return [...prev, { key: '', value: '' }];
        }
        return [{ key: '', value: '' }];
      });
    }
  }

  /** Removes the current key and value object from the tags array, but only if there is at least another key and value object in the tags array */
  function removeTag(id: number) {
    if (tags) {
      const newValues: Tag[] = [...tags.slice(0, id), ...tags.slice(id + 1)];
      setTags(newValues);
    }
  }

  /** Handles the tags saving/updating/deleting operations */
  function updateTags(action?: 'delete') {
    if (tags && data) {
      const serviceId = data.id;
      let payload;

      if (!action) {
        setLoading(true);
        payload = JSON.stringify(tags);
      } else {
        setDeleteLoading(true);
        payload = JSON.stringify([]);
      }

      settingsService.saveTags(serviceId, payload).then(res => {
        if (res === Error) {
          setLoading(false);
          setDeleteLoading(false);
          showToast({
            hasError: true,
            title: `Tags were not ${!action ? 'saved' : 'deleted'}!`,
            message: `There was an error ${
              !action ? 'saving' : 'deleting'
            } the tags. Please try again later.`
          });
        } else {
          setLoading(false);
          setDeleteLoading(false);
          showToast({
            hasError: false,
            title: `Tags have been ${!action ? 'saved' : 'deleted'}!`,
            message: `The tags have been ${!action ? 'saved' : 'deleted'} for ${
              data.provider
            } ${data.service} ${data.name}`
          });
          setSkipped(0);
          setInventoryHasUpdate(true);
          closeModal();
        }
      });
    }
  }

  /** Handles the bulk tags saving/deleting operations */
  function updateBulkTags(action?: 'delete') {
    if (!data && tags && bulkItems) {
      const payload = {
        resources: bulkItems,
        tags
      };

      if (!action) {
        setLoading(true);
      } else {
        setDeleteLoading(true);
        payload.tags = [];
      }

      const payloadJSON = JSON.stringify(payload);

      settingsService.bulkSaveTags(payloadJSON).then(res => {
        if (res === Error) {
          setLoading(false);
          setDeleteLoading(false);
          showToast({
            hasError: true,
            title: `Tags were not ${!action ? 'saved' : 'deleted'}!`,
            message: `There was an error ${
              !action ? 'saving' : 'deleting'
            } the tags. Please try again later.`
          });
        } else {
          setLoading(false);
          setDeleteLoading(false);
          showToast({
            hasError: false,
            title: `Tags have been ${!action ? 'saved' : 'deleted'}!`,
            message: `The tags have been ${!action ? 'saved' : 'deleted'} for ${
              bulkItems.length
            } ${bulkItems.length > 1 ? 'resources' : 'resource'}`
          });
          setInventoryHasUpdate(true);
          setBulkItems([]);
          closeModal();
        }
      });
    }
  }

  /** Handles the checkbox change for bulk actions, when ticking the checkbox resource by resource */
  function onCheckboxChange(e: ChangeEvent<HTMLInputElement>, id: string) {
    if (e.target.checked) {
      const newArray = [...bulkItems];
      newArray.push(id);
      setBulkItems(newArray);
    } else {
      const newArray = bulkItems.filter(currentId => currentId !== id);
      setBulkItems(newArray);
    }
  }

  /** Handles the checkbox change for all resources */
  function handleBulkSelection(e: ChangeEvent<HTMLInputElement>) {
    if (inventory && e.target.checked && !query) {
      const arrayOfIds = inventory.map(item => item.id);
      setBulkItems(arrayOfIds);
      setBulkSelectCheckbox(true);
    }

    if (inventory && !e.target.checked && !query) {
      setBulkItems([]);
      setBulkSelectCheckbox(false);
    }

    if (searchedInventory && e.target.checked) {
      const arrayOfIds = searchedInventory.map(item => item.id);
      setBulkItems(arrayOfIds);
      setBulkSelectCheckbox(true);
    }

    if (searchedInventory && !e.target.checked) {
      setBulkItems([]);
      setBulkSelectCheckbox(false);
    }
  }

  /** Handles the add hide resource operation from the bulk actions bottom bar */
  function hideResourceFromCustomView() {
    if (!router.query.view || bulkItems.length === 0) return;

    const currentHiddenResources: number[] = hiddenResources!.map(
      resource => resource.id
    );

    const newHiddenResources = [...currentHiddenResources, ...bulkItems];

    setHideResourcesLoading(true);

    const viewId = router.query.view.toString();
    const newPayload = { id: Number(viewId), exclude: newHiddenResources };
    const payload = JSON.stringify(newPayload);

    settingsService.hideResourceFromView(viewId, payload).then(res => {
      if (res === Error) {
        setHideResourcesLoading(false);
        showToast({
          hasError: true,
          title: 'Resources could not be hid.',
          message:
            'There was an error hiding the resources. Please refer to the logs and try again.'
        });
      } else {
        setHideResourcesLoading(false);
        showToast({
          hasError: false,
          title: 'Resources are now hidden.',
          message:
            'The resources were successfully hidden. They can be unhid from the custom view management.'
        });
        setHideOrUnhideHasUpdate(true);
      }
    });
  }

  /** Handles the filter removal from the InventorySummary and also the URL params */
  function deleteFilter(idx: number) {
    const updatedFilters: InventoryFilterData[] = [...filters!];
    updatedFilters.splice(idx, 1);
    const url = updatedFilters
      .map(
        filter =>
          `${filter.field}${`:${filter.operator}`}${
            filter.values.length > 0 ? `:${filter.values}` : ''
          }`
      )
      .join('&');
    router.push(url ? `?${url}` : '', undefined, { shallow: true });
  }

  /** Conditional helpers */
  const hasErrorAndNoInventory =
    (error && !inventoryStats) || (error && !inventory);

  const hasNoInventory =
    (!error && inventoryStats && Object.keys(inventoryStats).length === 0) ||
    (!error && inventory && inventory.length === 0);

  const isNotCustomView = !router.query.view;

  const hasFilterOrCustomView =
    (filters && filters.length > 0) || router.query.view;

  const displayFilterIfIsNotCustomView =
    !error &&
    !router.query.view &&
    ((inventory && inventory.length > 0) ||
      (searchedInventory && searchedInventory.length > 0));

  const loadingFilters =
    Object.keys(router.query).length > 0 && !displayedFilters && !error;

  const hasFilters =
    Object.keys(router.query).length > 0 &&
    displayedFilters &&
    displayedFilters.length > 0;

  const loadingInventory =
    !inventory && !error && !query && !displayedFilters && !router.query.view;

  return {
    inventoryStats,
    inventory,
    searchedInventory,
    error,
    query,
    setQuery,
    openModal,
    isOpen,
    closeModal,
    data,
    page,
    goTo,
    tags,
    handleChange,
    addNewTag,
    removeTag,
    loading,
    updateTags,
    toast,
    showToast,
    dismissToast,
    deleteLoading,
    reloadDiv,
    bulkItems,
    onCheckboxChange,
    handleBulkSelection,
    bulkSelectCheckbox,
    openBulkModal,
    updateBulkTags,
    router,
    filters,
    displayedFilters,
    setSkippedSearch,
    deleteFilter,
    searchedLoading,
    statsLoading,
    views,
    getViews,
    hiddenResources,
    hideResourceFromCustomView,
    hideResourcesLoading,
    setHideOrUnhideHasUpdate,
    hasErrorAndNoInventory,
    hasNoInventory,
    isNotCustomView,
    hasFilterOrCustomView,
    displayFilterIfIsNotCustomView,
    loadingFilters,
    hasFilters,
    loadingInventory,
    isSomeServiceUnavailable
  };
}

export default useInventory;
