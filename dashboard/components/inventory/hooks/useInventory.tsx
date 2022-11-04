import { ChangeEvent, RefObject, useEffect, useState } from "react";
import settingsService from "../../../services/settingsService";
import { Provider } from "../../../utils/providerHelper";
import useToast from "../../toast/hooks/useToast";
import useIsVisible from "./useOnScreen";

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

type Pages = "tags" | "delete";

function useInventory(reloadDiv: RefObject<HTMLDivElement>) {
  const [inventoryStats, setInventoryStats] = useState<
    InventoryStats | undefined
  >();
  const [inventory, setInventory] = useState<InventoryItem[] | undefined>();
  const [error, setError] = useState(false);
  const [skipped, setSkipped] = useState(0);
  const [inventoryHasUpdate, setInventoryHasUpdate] = useState(false);
  const [query, setQuery] = useState("");
  const [searchedInventory, setSearchedInventory] = useState<
    InventoryItem[] | undefined
  >();
  const [isOpen, setIsOpen] = useState(false);
  const [data, setData] = useState<InventoryItem>();
  const [page, setPage] = useState<Pages>("tags");
  const [tags, setTags] = useState<Tag[]>();
  const [loading, setLoading] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);

  const { toast, setToast, dismissToast } = useToast();

  const isVisible = useIsVisible(reloadDiv);

  // First fetch effect
  useEffect(() => {
    let mounted = true;

    setError(false);

    settingsService.getInventoryStats().then((res) => {
      if (mounted) {
        if (res === Error) {
          setError(true);
        } else {
          setInventoryStats(res);
        }
      }
    });

    settingsService
      .getInventoryList(`?limit=50&skip=${skipped}`)
      .then((res) => {
        if (mounted) {
          if (res === Error) {
            setError(true);
          } else {
            setInventory((prev) => {
              if (prev) {
                return [...prev, ...res];
              }
              return res;
            });
            setSkipped((prev) => prev + 50);
          }
        }
      });

    return () => {
      mounted = false;
    };
  }, []);

  // Infinite scrolling fetch effect
  /* useEffect(() => {
    let mounted = true;

    if (
      workspaceId &&
      inventoryStats &&
      skipped < inventoryStats.resources &&
      isVisible &&
      !query
    ) {
      setError(false);

      settingsService
        .getInventoryList(workspaceId, `?limit=50&skip=${skipped}`)
        .then((res) => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setInventory((prev) => {
                if (prev) {
                  return [...prev, ...res];
                }
                return res;
              });
              setSkipped((prev) => prev + 50);
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

    setSearchedInventory(undefined);

    if (query && workspaceId) {
      setError(false);

      setTimeout(() => {
        if (mounted) {
          settingsService.searchInventory(workspaceId, query).then((res) => {
            if (mounted) {
              if (res === Error) {
                setError(true);
              } else {
                setSearchedInventory(res);
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

    if (workspaceId && inventoryHasUpdate) {
      settingsService
        .getInventoryList(workspaceId, `?limit=50&skip=0`)
        .then((res) => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setQuery("");
              setInventory(res);
              setSkipped(50);
              setInventoryHasUpdate(false);
            }
          }
        });
    }

    return () => {
      mounted = false;
    };
  }, [inventoryHasUpdate]); */

  // Listen to ESC key on modal effect
  useEffect(() => {
    function escFunction(event: KeyboardEvent) {
      if (event.key === "Escape") {
        setIsOpen(false);
      }
    }

    document.addEventListener("keydown", escFunction, false);

    return () => {
      document.removeEventListener("keydown", escFunction, false);
    };
  }, []);

  // Functions to be exported
  function cleanModal() {
    setData(undefined);
    setPage("tags");
  }

  function openModal(inventoryItem: InventoryItem) {
    cleanModal();

    setData(inventoryItem);

    if (inventoryItem.tags && inventoryItem.tags.length > 0) {
      setTags(inventoryItem.tags);
    } else {
      setTags([{ key: "", value: "" }]);
    }

    setIsOpen(true);
  }

  function closeModal() {
    setIsOpen(false);
  }

  function goTo(newPage: Pages) {
    setPage(newPage);
  }

  function handleChange(newData: Partial<Tag>, id?: number) {
    if (tags && typeof id === "number") {
      const newValues: Tag[] = [...tags];
      newValues[id] = {
        ...newValues[id],
        ...newData,
      };
      setTags(newValues);
    }
  }

  function addNewTag() {
    if (tags) {
      setTags((prev) => {
        if (prev) {
          return [...prev, { key: "", value: "" }];
        }
        return [{ key: "", value: "" }];
      });
    }
  }

  function removeTag(id: number) {
    if (tags) {
      const newValues: Tag[] = [...tags.slice(0, id), ...tags.slice(id + 1)];
      setTags(newValues);
    }
  }

  function updateTags(action?: "delete") {
    /* if (tags && workspaceId && data) {
      const serviceId = data.id;
      let payload;

      if (!action) {
        setLoading(true);
        payload = JSON.stringify(tags);
      } else {
        setDeleteLoading(true);
        payload = JSON.stringify([]);
      }

      settingsService.saveTags(workspaceId, serviceId, payload).then((res) => {
        if (res === Error) {
          setLoading(false);
          setDeleteLoading(false);
          setToast({
            hasError: true,
            title: `Tags were not ${!action ? "saved" : "deleted"}!`,
            message: `There was an error ${
              !action ? "saving" : "deleting"
            } the tags. Please try again later.`,
          });
        } else {
          setLoading(false);
          setDeleteLoading(false);
          setToast({
            hasError: false,
            title: `Tags have been ${!action ? "saved" : "deleted"}!`,
            message: `The tags have been ${!action ? "saved" : "deleted"} for ${
              data.cloud
            } ${data.service} ${data.name}`,
          });
          setInventoryHasUpdate(true);
          closeModal();
        }
      });
    } */
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
  };
}

export default useInventory;
