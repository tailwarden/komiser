import { useRouter } from 'next/router';
import { ChangeEvent, useEffect, useRef, useState } from 'react';
import settingsService from '../../../services/settingsService';
import { Provider } from '../../../utils/providerHelper';
import useToast from '../../toast/hooks/useToast';
import useIsVisible from './useIsVisible';

export type InventoryFilterDataProps = {
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

export type ViewProps = {
  id: number;
  name: string;
  filters: InventoryFilterDataProps[];
  exclude: string[];
};

function useInventory() {
  const [inventoryStats, setInventoryStats] = useState<
    InventoryStats | undefined
  >();
  const [inventory, setInventory] = useState<InventoryItem[] | undefined>();
  const [error, setError] = useState(false);
  const [skipped, setSkipped] = useState(0);
  const [skippedSearch, setSkippedSearch] = useState(0);
  const [inventoryHasUpdate, setInventoryHasUpdate] = useState(false);
  const [query, setQuery] = useState('');
  const [searchedInventory, setSearchedInventory] = useState<
    InventoryItem[] | undefined
  >();
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
    useState<InventoryFilterDataProps[]>();
  const [filters, setFilters] = useState<InventoryFilterDataProps[]>();
  const [searchedLoading, setSearchedLoading] = useState(false);
  const [statsLoading, setStatsLoading] = useState(false);
  const [views, setViews] = useState<ViewProps[]>();

  const { toast, setToast, dismissToast } = useToast();
  const reloadDiv = useRef<HTMLDivElement>(null);
  const isVisible = useIsVisible(reloadDiv);
  const batchSize: number = 50;
  const router = useRouter();

  // Parse URL params
  function parseParams(
    param: string | InventoryFilterDataProps,
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

    return filter as InventoryFilterDataProps;
  }

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

  function getViews(edit?: boolean, viewId?: string) {
    settingsService.getViews().then(res => {
      if (res === Error) {
        setToast({
          hasError: true,
          title: `The custom views couldn't be loaded.`,
          message: `There was a problem fetching the views. Please try again.`
        });
      } else {
        const sortViewsById: ViewProps[] = res;
        sortViewsById.sort((a, b) => a.id - b.id);
        setViews(sortViewsById);
        if (edit && viewId) {
          router.push(`/?view=${viewId}`, undefined, { shallow: true });
        }
      }
    });
  }

  // Getting all the views
  useEffect(() => {
    getViews();
  }, []);

  // Fetch the correct inventory list based on URL params
  useEffect(() => {
    let mounted = true;
    resetStates();

    // Fetch base inventory
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
        if (mounted) {
          if (res === Error) {
            setError(true);
            setStatsLoading(false);
          } else {
            setInventoryStats(res);
            setStatsLoading(false);
          }
        }
      });

      settingsService
        .getInventoryList(`?limit=${batchSize}&skip=${skipped}`)
        .then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setInventory(res);
              setSkipped(prev => prev + batchSize);
            }
          }
        });
    }

    // Fetch from a custom view
    if (router.query.view && views && views.length > 0) {
      const filterFound = views.find(
        view => view.id.toString() === router.query.view
      );

      if (filterFound) {
        setSearchedLoading(true);
        setStatsLoading(true);
        const payloadJson = JSON.stringify(filterFound?.filters);

        settingsService.getFilteredInventoryStats(payloadJson).then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
              setStatsLoading(false);
            } else {
              setInventoryStats(res);
              setStatsLoading(false);
            }
          }
        });

        settingsService
          .getFilteredInventory(`?limit=${batchSize}&skip=0`, payloadJson)
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
                const newFiltersToDisplay: InventoryFilterDataProps[] =
                  filterFound!.filters.map(filter =>
                    parseParams(filter, 'display', true)
                  );
                setDisplayedFilters(newFiltersToDisplay);

                if (res.length >= batchSize) {
                  setShouldFetchMore(true);
                } else {
                  setShouldFetchMore(false);
                }
              }
            }
          });
      }
    }

    // Fetch from filters
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

      const newFilters: InventoryFilterDataProps[] = Object.keys(
        router.query
      ).map(param => parseParams(param as string, 'fetch'));
      const newFiltersToDisplay: InventoryFilterDataProps[] = Object.keys(
        router.query
      ).map(param => parseParams(param as string, 'display'));

      const payloadJson = JSON.stringify(newFilters);

      settingsService.getFilteredInventoryStats(payloadJson).then(res => {
        if (mounted) {
          if (res === Error) {
            setError(true);
            setStatsLoading(false);
          } else {
            setInventoryStats(res);
            setStatsLoading(false);
          }
        }
      });

      settingsService
        .getFilteredInventory(`?limit=${batchSize}&skip=0`, payloadJson)
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
          }
        });
    }

    return () => {
      mounted = false;
    };
  }, [router.query, views]);

  // Infinite scrolling fetch effect
  useEffect(() => {
    let mounted = true;

    // Infinite scrolling fetch on normal list
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
          if (mounted) {
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
          }
        });
    }

    // Infinite scrolling fetch on searched normal inventory
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
          if (mounted) {
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
          }
        });
    }

    // Infinite scrolling fetch on searched filtered list or custom view
    if (
      shouldFetchMore &&
      isVisible &&
      query &&
      Object.keys(router.query).length > 0
    ) {
      let payloadJson = '';

      if (!router.query.view && filters && filters.length > 0) {
        payloadJson = JSON.stringify(filters);
      }

      if (router.query.view && views && views.length > 0) {
        const filterFound = views.find(
          view => view.id.toString() === router.query.view
        );
        payloadJson = JSON.stringify(filterFound?.filters);
      }

      settingsService
        .getFilteredInventory(
          `?limit=${batchSize}&skip=${skippedSearch}&query=${query}`,
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
          }
        });
    }

    // Infinite scrolling fetch on filtered list
    if (shouldFetchMore && isVisible && filters && !query) {
      setError(false);

      const payloadJson = JSON.stringify(filters);
      settingsService
        .getFilteredInventory(
          `?limit=${batchSize}&skip=${skippedSearch}`,
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
          }
        });
    }

    // Infinite scrolling fetch on custom view
    if (
      shouldFetchMore &&
      isVisible &&
      views &&
      views.length > 0 &&
      router.query.view &&
      !query
    ) {
      const filterFound = views.find(
        view => view.id.toString() === router.query.view
      );

      if (filterFound) {
        const payloadJson = JSON.stringify(filterFound?.filters);

        settingsService
          .getFilteredInventory(
            `?limit=${batchSize}&skip=${skippedSearch}`,
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
            }
          });
      }
    }

    return () => {
      mounted = false;
    };
  }, [isVisible]);

  // Search effect behavior
  // If there's a filtered list, search should only bring back results from the list
  // If not, search should get from all inventory
  // A filter can overwrite a search, but not the opposite
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

  // Tags saved list refresh effect
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

  // Functions to be exported
  function cleanModal() {
    setData(undefined);
    setPage('tags');
  }

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

  function openBulkModal() {
    cleanModal();
    setTags([{ key: '', value: '' }]);
    setIsOpen(true);
  }

  function closeModal() {
    setIsOpen(false);
  }

  function goTo(newPage: Pages) {
    setPage(newPage);
  }

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

  function removeTag(id: number) {
    if (tags) {
      const newValues: Tag[] = [...tags.slice(0, id), ...tags.slice(id + 1)];
      setTags(newValues);
    }
  }

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

  function deleteFilter(idx: number) {
    const updatedFilters: InventoryFilterDataProps[] = [...filters!];
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
    getViews
  };
}

export default useInventory;
