import { useRouter } from 'next/router';
import { ChangeEvent, useEffect, useRef, useState } from 'react';
import settingsService from '../../../services/settingsService';
import { Provider } from '../../../utils/providerHelper';
import useToast from '../../toast/hooks/useToast';
import useIsVisible from './useIsVisible';

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
export type Pages = 'tags' | 'delete';

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

function useInventory() {
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
  const [page, setPage] = useState<Pages>('tags');
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

  const { toast, setToast, dismissToast } = useToast();
  const reloadDiv = useRef<HTMLDivElement>(null);
  const isVisible = useIsVisible(reloadDiv);
  const batchSize: number = 50;
  const router = useRouter();

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

  /** Parse the URL Params.
   * - Argument of type 'fetch' will output the object to fetch an inventory list and stats based on filters.
   * - Argument of type 'display' will output the object to populate the InventoryFilterSummary component */
  function parseURLParams(
    param: string | InventoryFilterData,
    type: 'fetch' | 'display',
    view?: boolean
  ) {
    let formatString;
    let filter;

    if (!view) {
      formatString = (param as string).split(':');
    } else {
      formatString = Object.values(param);
      formatString = [...formatString.slice(0, 2), formatString[2]!.toString()];
    }

    if (formatString[0]!.includes('tag:')) {
      const tag = (formatString[0] as string).split(':');
      formatString = [
        tag[0],
        tag[1],
        formatString[1],
        formatString[2]?.toString()
      ];
    }

    if (formatString[0] === 'tag' && type === 'fetch') {
      if (formatString.length > 2) {
        if (
          formatString.indexOf('IS_EMPTY') !== -1 ||
          formatString.indexOf('IS_NOT_EMPTY') !== -1
        ) {
          const key = formatString.slice(1, formatString.length - 1).join(':');

          filter = {
            field: `${formatString[0]}:${key}`,
            operator: formatString[formatString.length - 1],
            values: []
          };
        } else {
          const key = formatString.slice(1, formatString.length - 2).join(':');

          filter = {
            field: `${formatString[0]}:${key}`,
            operator: formatString[formatString.length - 2],
            values: (formatString[formatString.length - 1] as string).split(',')
          };
        }
      } else {
        filter = {
          field: `${formatString[0]}:${formatString[1]}`,
          operator: formatString[2],
          values:
            formatString[2] === 'IS_EMPTY' || formatString[2] === 'IS_NOT_EMPTY'
              ? []
              : (formatString[3] as string).split(',')
        };
      }
    }

    if (formatString[0] !== 'tag' && type === 'fetch') {
      filter = {
        field: formatString[0],
        operator: formatString[1],
        values:
          formatString[1] === 'IS_EMPTY' || formatString[1] === 'IS_NOT_EMPTY'
            ? []
            : (formatString[2] as string).split(',')
      };
    }

    if (formatString[0] === 'tag' && type === 'display') {
      if (formatString.length > 2) {
        if (
          formatString.indexOf('IS_EMPTY') !== -1 ||
          formatString.indexOf('IS_NOT_EMPTY') !== -1
        ) {
          const key = formatString.slice(1, formatString.length - 1).join(':');

          filter = {
            field: formatString[0],
            tagKey: key,
            operator: formatString[formatString.length - 1],
            values: []
          };
        } else {
          const key = formatString.slice(1, formatString.length - 2).join(':');

          filter = {
            field: formatString[0],
            tagKey: key,
            operator: formatString[formatString.length - 2],
            values: (formatString[formatString.length - 1] as string).split(',')
          };
        }
      } else {
        filter = {
          field: formatString[0],
          tagKey: formatString[1],
          operator: formatString[2],
          values:
            formatString[2] === 'IS_EMPTY' || formatString[2] === 'IS_NOT_EMPTY'
              ? []
              : (formatString[3] as string).split(',')
        };
      }
    }

    if (formatString[0] !== 'tag' && type === 'display') {
      filter = {
        field: formatString[0],
        operator: formatString[1],
        values:
          formatString[1] === 'IS_EMPTY' || formatString[1] === 'IS_NOT_EMPTY'
            ? []
            : (formatString[2] as string).split(',')
      };
    }

    return filter as InventoryFilterData;
  }

  /** Fetch all the custom views.
   * - Views will be stored in the state: views.
   */
  function getViews(edit?: boolean, viewId?: string, redirect?: boolean) {
    settingsService.getViews().then(res => {
      if (res === Error) {
        setToast({
          hasError: true,
          title: `The custom views couldn't be loaded.`,
          message: `There was a problem fetching the views. Please try again.`
        });
      } else {
        const sortViewsById: View[] = res;
        sortViewsById.sort((a, b) => a.id - b.id);
        setViews(sortViewsById);
        if (edit && viewId) {
          router.push(`/?view=${viewId}`, undefined, { shallow: true });
        }

        if (redirect) {
          const idFromNewCustomView = res[res.length - 1].id;
          router.push(`/?view=${idFromNewCustomView}`, undefined, {
            shallow: true
          });
        }
      }
    });
  }

  /** Fetch base inventory and top stats for All Resources.
   * - Inventory list will be stored in the state: inventory
   * - Inventory top stats will be stored in the state: inventoryStats
   */
  function getInventoryListAndStats() {
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

      settingsService
        .getInventoryList(`?limit=${batchSize}&skip=${skipped}`)
        .then(res => {
          if (res === Error) {
            setError(true);
          } else {
            setInventory(res);
            setSkipped(prev => prev + batchSize);
          }
        });
    }
  }

  /** Fetch inventory from a filter.
   * - Inventory list will be stored in the state: searchedInventory
   */
  function getInventoryListFromAFilter() {
    if (
      router.query &&
      Object.keys(router.query).length > 0 &&
      !router.query.view
    ) {
      if (
        Object.keys(router.query)[0].split(':').length <= 1 &&
        !router.query.view
      ) {
        setTimeout(() => router.push('/'), 5000);
        return setToast({
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
        .getFilteredInventory(`?limit=${batchSize}&skip=0`, payloadJson)
        .then(res => {
          if (res.error) {
            setToast({
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

  /** Fetch inventory, top stats and hidden resources for a given custom view.
   * - Inventory list will be stored in the state: searchedInventory
   * - Inventory top stats will be stored in the state: inventoryStats
   */
  function getCustomViewInventoryListAndStats() {
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
          .getCustomViewInventory(
            id as string,
            `?limit=${batchSize}&skip=0`,
            payloadJson
          )
          .then(res => {
            if (res.error) {
              setToast({
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
      } else {
        setTimeout(() => router.push('/'), 5000);
        return setToast({
          hasError: true,
          title: `Invalid view`,
          message: `We couldn't find this view. Redirecting back to home...`
        });
      }
    }
    return null;
  }

  /** Whenever the page changes or views is updated, fetch the custom views, reset UI states and fetch the inventory based on the URL params */
  useEffect(() => {
    if (!views) {
      getViews();
    }
    resetStates();
    getInventoryListAndStats();
    getCustomViewInventoryListAndStats();
    getInventoryListFromAFilter();
  }, [router.query, views]);

  /** When a resource is hid or unhid, reset the states and call getCustomViewInventoryListAndStats. */
  useEffect(() => {
    if (hideOrUnhideHasUpdate) {
      resetStates();
      getCustomViewInventoryListAndStats();
    }
  }, [hideOrUnhideHasUpdate]);

  /** Load the next 50 results when the user scrolls the inventory list to the end */
  function infiniteScrollInventoryList() {
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
        .getInventoryList(`?limit=${batchSize}&skip=${skipped}`)
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

  /** Load the next 50 results when the user scrolls a filtered inventory list to the end */
  function infiniteScrollFilteredList() {
    if (shouldFetchMore && isVisible && filters && !query) {
      setError(false);

      const payloadJson = JSON.stringify(filters);
      settingsService
        .getFilteredInventory(
          `?limit=${batchSize}&skip=${skippedSearch}`,
          payloadJson
        )
        .then(res => {
          if (res.error) {
            setToast({
              hasError: true,
              title: `Filter could not be applied!`,
              message: `Please refresh the page and try again.`
            });
            setLoading(false);
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

  /** Load the next 50 results when the user scrolls a searched inventory list to the end */
  function infiniteScrollSearchedList() {
    if (
      shouldFetchMore &&
      isVisible &&
      query &&
      Object.keys(router.query).length === 0
    ) {
      setError(false);

      settingsService
        .getInventoryList(
          `?limit=${batchSize}&skip=${skippedSearch}&query=${query}`
        )
        .then(res => {
          if (res === Error) {
            setError(true);
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

  /** Load the next 50 results when the user scrolls a searched and filtered inventory list to the end */
  function infiniteScrollSearchedAndFilteredList() {
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
        .getFilteredInventory(
          `?limit=${batchSize}&skip=${skippedSearch}&query=${query}`,
          payloadJson
        )
        .then(res => {
          if (res.error) {
            setToast({
              hasError: true,
              title: `Filter could not be applied!`,
              message: `Please refresh the page and try again.`
            });
            setLoading(false);
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

  /** Load the next 50 results when the user scrolls a custom view list to the end */
  function infiniteScrollCustomViewList() {
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
          .getCustomViewInventory(
            id as string,
            `?limit=${batchSize}&skip=${skippedSearch}`,
            payloadJson
          )
          .then(res => {
            if (res.error) {
              setToast({
                hasError: true,
                title: `Filter could not be applied!`,
                message: `Please refresh the page and try again.`
              });
              setError(true);
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

  /** Load the next 50 results when the user scrolls a searched custom view list to the end */
  function infiniteScrollSearchedCustomViewList() {
    if (
      shouldFetchMore &&
      isVisible &&
      query &&
      router.query.view &&
      views &&
      views.length > 0
    ) {
      const id = router.query.view;
      const filterFound = views.find(view => view.id.toString() === id);

      if (filterFound) {
        const payloadJson = JSON.stringify(filterFound?.filters);

        settingsService
          .getCustomViewInventory(
            id as string,
            `?limit=${batchSize}&skip=${skippedSearch}`,
            payloadJson
          )
          .then(res => {
            if (res.error) {
              setToast({
                hasError: true,
                title: `Filter could not be applied!`,
                message: `Please refresh the page and try again.`
              });
              setError(true);
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

  /** Infinite scrolling handler. Identifies which inventory should be fetched on scroll. */
  useEffect(() => {
    infiniteScrollInventoryList();
    infiniteScrollFilteredList();
    infiniteScrollSearchedList();
    infiniteScrollSearchedAndFilteredList();
    infiniteScrollCustomViewList();
    infiniteScrollSearchedCustomViewList();
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

    if (!filters && !query && !router.query.view) {
      setSearchedInventory(undefined);
    }

    if (!filters && query && !router.query.view) {
      setSearchedLoading(true);
      setError(false);
      setTimeout(() => {
        if (mounted) {
          settingsService
            .getInventoryList(
              `?limit=${batchSize}&skip=0${query && `&query=${query}`}`
            )
            .then(res => {
              if (mounted) {
                if (res === Error) {
                  setError(true);
                  setSearchedLoading(false);
                }
                setSearchedInventory(res);
                setSearchedLoading(false);

                if (res.length >= batchSize) {
                  setShouldFetchMore(true);
                  setSkippedSearch(prev => prev + batchSize);
                }
              }
            });
        }
      }, 700);
    }

    if (filters && filters.length > 0 && !router.query.view) {
      const payloadJson = JSON.stringify(filters);
      setSearchedLoading(true);
      setTimeout(() => {
        if (mounted) {
          settingsService
            .getFilteredInventory(
              `?limit=${batchSize}&skip=0${query && `&query=${query}`}`,
              payloadJson
            )
            .then(res => {
              if (mounted) {
                if (res.error) {
                  setToast({
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

    if (router.query.view && views && views.length > 0) {
      const filterFound = views.find(
        view => view.id.toString() === router.query.view
      );

      if (filterFound) {
        const payloadJson = JSON.stringify(filterFound.filters);
        setSearchedLoading(true);
        setTimeout(() => {
          if (mounted) {
            settingsService
              .getFilteredInventory(
                `?limit=${batchSize}&skip=0${query && `&query=${query}`}`,
                payloadJson
              )
              .then(res => {
                if (mounted) {
                  if (res.error) {
                    setToast({
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
    }
    return () => {
      mounted = false;
    };
  }, [query]);

  /** Refresh list when tags are saved.
   * - If it's on all resources, refetch the inventory list
   * - If it's on a custom view, quick reload the current url
   */
  useEffect(() => {
    let mounted = true;

    if (inventoryHasUpdate) {
      if (Object.keys(router.query).length === 0) {
        settingsService
          .getInventoryList(`?limit=${batchSize}&skip=0`)
          .then(res => {
            if (mounted) {
              if (res === Error) {
                setError(true);
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
    setPage('tags');
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
          setToast({
            hasError: true,
            title: `Tags were not ${!action ? 'saved' : 'deleted'}!`,
            message: `There was an error ${
              !action ? 'saving' : 'deleting'
            } the tags. Please try again later.`
          });
        } else {
          setLoading(false);
          setDeleteLoading(false);
          setToast({
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
          setToast({
            hasError: true,
            title: `Tags were not ${!action ? 'saved' : 'deleted'}!`,
            message: `There was an error ${
              !action ? 'saving' : 'deleting'
            } the tags. Please try again later.`
          });
        } else {
          setLoading(false);
          setDeleteLoading(false);
          setToast({
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
        setToast({
          hasError: true,
          title: 'Resources could not be hid.',
          message:
            'There was an error hiding the resources. Please refer to the logs and try again.'
        });
      } else {
        setHideResourcesLoading(false);
        setToast({
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
    router.push(url ? `/?${url}` : '', undefined, { shallow: true });
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
    Object.keys(router.query).length > 0 && !displayedFilters;

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
    setToast,
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
    loadingInventory
  };
}

export default useInventory;
