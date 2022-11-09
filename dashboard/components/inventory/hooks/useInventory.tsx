import { ChangeEvent, useEffect, useRef, useState } from 'react';
import settingsService from '../../../services/settingsService';
import { Provider } from '../../../utils/providerHelper';
import useToast from '../../toast/hooks/useToast';
import useIsVisible from './useOnScreen';

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
  provider: Provider;
  cost: number;
  id: string;
  name: string;
  region: string;
  service: string;
  tags: Tag[] | [] | null;
  workspaceId: string;
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

  const { toast, setToast, dismissToast } = useToast();
  const reloadDiv = useRef<HTMLDivElement>(null);
  const isVisible = useIsVisible(reloadDiv);
  const batchSize: number = 50;

  // First fetch effect
  useEffect(() => {
    let mounted = true;

    setError(false);

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

    return () => {
      mounted = false;
    };
  }, []);

  // Infinite scrolling fetch effect
  useEffect(() => {
    let mounted = true;

    // Fetching on unsearched list
    if (
      inventoryStats &&
      skipped < inventoryStats.resources &&
      isVisible &&
      !query
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

    return () => {
      mounted = false;
    };
  }, [isVisible]);

  // Search effect
  useEffect(() => {
    let mounted = true;
    setSearchedInventory(undefined);
    setSkippedSearch(0);
    setShouldFetchMore(false);
    setBulkItems([]);
    setBulkSelectCheckbox(false);

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

      console.log(payload);
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
            } ${bulkItems.length > 1 ? 'items' : 'item'}`
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

    if (searchedInventory && e.target.checked && query) {
      const arrayOfIds = searchedInventory.map(item => item.id);
      setBulkItems(arrayOfIds);
      setBulkSelectCheckbox(true);
    }

    if (searchedInventory && !e.target.checked && query) {
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
    dismissToast,
    deleteLoading,
    reloadDiv,
    bulkItems,
    onCheckboxChange,
    handleBulkSelection,
    bulkSelectCheckbox,
    openBulkModal,
    updateBulkTags
  };
}

export default useInventory;
