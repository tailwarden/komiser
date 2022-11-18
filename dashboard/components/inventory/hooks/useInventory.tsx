import { useRouter } from 'next/router';
import { ChangeEvent, useEffect, useRef, useState } from 'react';
import settingsService from '../../../services/settingsService';
import { Provider } from '../../../utils/providerHelper';
import useToast from '../../toast/hooks/useToast';
import useIsVisible from './useOnScreen';

export type InventoryFilterDataProps = {
  field:
    | 'provider'
    | 'region'
    | 'account'
    | 'name'
    | 'service'
    | string
    | undefined;
  operator:
    | 'IS'
    | 'IS_NOT'
    | 'CONTAINS'
    | 'NOT_CONTAINS'
    | 'IS_EMPTY'
    | 'IS_NOT_EMPTY'
    | undefined;
  tagKey: string | undefined;
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
  const [activeFilters, setActiveFilters] =
    useState<InventoryFilterDataProps>();

  const { toast, setToast, dismissToast } = useToast();
  const reloadDiv = useRef<HTMLDivElement>(null);
  const isVisible = useIsVisible(reloadDiv);
  const batchSize: number = 50;
  const router = useRouter();

  function resetStates() {
    setSkipped(0);
    setInventory(undefined);
    setInventoryStats(undefined);
    setSearchedInventory(undefined);
    setActiveFilters(undefined);
    setBulkItems([]);
    setBulkSelectCheckbox(false);
    setQuery('');
  }

  // Fetch the correct inventory list based on URL params
  useEffect(() => {
    let mounted = true;
    resetStates();

    if (router.query && Object.keys(router.query).length === 0) {
      settingsService.getInventoryStats().then(res => {
        if (mounted) {
          if (res === Error) {
            setError(true);
          } else {
            setInventoryStats(res);
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

    if (router.query.field) {
      const formatValue =
        router.query.values && (router.query.values as string).split(',');
      const isTagFilter =
        router.query.field && router.query.field.includes('tag');

      const filterProps: any = {
        field: isTagFilter ? 'tag' : router.query.field,
        operator: router.query.operator,
        values: formatValue,
        tagKey: isTagFilter ? router.query.field!.slice(4) : ''
      };

      const payload = [
        {
          field: router.query.field,
          operator: router.query.operator,
          values: formatValue
        }
      ];

      const payloadJson = JSON.stringify(payload);

      settingsService.getFilteredInventoryStats(payloadJson).then(res => {
        if (mounted) {
          if (res === Error) {
            setError(true);
          } else {
            setInventoryStats(res);
          }
        }
      });

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
              setSearchedInventory(prev => {
                if (prev) {
                  return [...prev, ...res];
                }
                return res;
              });
              setActiveFilters(filterProps);
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

    return () => {
      mounted = false;
    };
  }, [router.query]);

  // Infinite scrolling fetch effect
  useEffect(() => {
    let mounted = true;

    // Fetching on unsearched list
    if (
      inventoryStats &&
      skipped < inventoryStats.resources &&
      isVisible &&
      !query &&
      !activeFilters
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

    // Fetching on searched list
    if (shouldFetchMore && isVisible && query) {
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

    // Fetching on filtered list
    if (shouldFetchMore && isVisible && activeFilters) {
      setError(false);
      const formatValue =
        router.query.values && (router.query.values as string).split(',');
      const isTagFilter =
        router.query.field && router.query.field.includes('tag');

      const filterProps: any = {
        field: isTagFilter ? 'tag' : router.query.field,
        operator: router.query.operator,
        values: formatValue,
        tagKey: isTagFilter ? router.query.field!.slice(4) : ''
      };

      const payload = [
        {
          field: router.query.field,
          operator: router.query.operator,
          values: formatValue
        }
      ];

      const payloadJson = JSON.stringify(payload);

      settingsService.getFilteredInventoryStats(payloadJson).then(res => {
        if (mounted) {
          if (res === Error) {
            setError(true);
          } else {
            setInventoryStats(res);
          }
        }
      });

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
              setActiveFilters(filterProps);

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

    return () => {
      mounted = false;
    };
  }, [isVisible]);

  // Search effect
  useEffect(() => {
    let mounted = true;

    if (!query) {
      setSearchedInventory(undefined);
      setSkippedSearch(0);
      setShouldFetchMore(false);
      setBulkItems([]);
      setBulkSelectCheckbox(false);
    }

    if (query) {
      setError(false);
      setTimeout(() => {
        if (mounted) {
          settingsService
            .getInventoryList(`?limit=${batchSize}&skip=0&query=${query}`)
            .then(res => {
              if (mounted) {
                if (res === Error) {
                  setError(true);
                }

                setSearchedInventory(res);

                if (res.length >= batchSize) {
                  setShouldFetchMore(true);
                  setSkippedSearch(prev => prev + batchSize);
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

  // Tags saved list refresh effect
  useEffect(() => {
    let mounted = true;

    if (inventoryHasUpdate) {
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
    }

    return () => {
      mounted = false;
    };
  }, [inventoryHasUpdate]);

  // Listen to ESC key on modal effect
  useEffect(() => {
    function escFunction(event: KeyboardEvent) {
      if (event.key === 'Escape') {
        setIsOpen(false);
      }
    }

    document.addEventListener('keydown', escFunction, false);

    return () => {
      document.removeEventListener('keydown', escFunction, false);
    };
  }, []);

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

  function openBulkModal(bulkItemsIds: string[]) {
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
    activeFilters,
    setSkippedSearch
  };
}

export default useInventory;
