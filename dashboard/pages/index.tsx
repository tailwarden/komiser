import Head from "next/head";
import { useRouter } from "next/router";
import { useRef } from "react";
import Button from "../components/button/Button";
import ErrorPage from "../components/error/ErrorPage";
import InventorySearchBar from "../components/inventory/components/InventorySearchBar";
import InventorySidePanel from "../components/inventory/components/InventorySidepanel";
import InventoryStatsCards from "../components/inventory/components/InventoryStatsCards";
import InventoryTable from "../components/inventory/components/InventoryTable";
import TagWrapper from "../components/inventory/components/TagWrapper";
import useInventory from "../components/inventory/hooks/useInventory";
import SkeletonInventory from "../components/skeleton/SkeletonInventory";
import SkeletonStats from "../components/skeleton/SkeletonStats";
import Toast from "../components/toast/Toast";
import formatNumber from "../utils/formatNumber";
import providers from "../utils/providerHelper";

export default function Inventory() {
  const reloadDiv = useRef<HTMLDivElement>(null);
  const router = useRouter();
  const {
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
  } = useInventory(reloadDiv);

  console.log(inventory);

  return (
    <>
      <Head>
        <title>Inventory - Komiser</title>
        <meta name="description" content="Inventory - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <p className="text-xl font-medium text-black-900">Inventory</p>
      <div className="mt-8"></div>

      {/* Toaster */}
      {toast && <Toast {...toast} dismissToast={dismissToast} />}

      {/* Error page */}
      {((error && !inventoryStats) || (error && !inventory)) && (
        <ErrorPage
          title="Network request error."
          message="There was an error fetching the inventory resources. Check out the server logs for more info and try again."
          action={
            <Button style="outline" onClick={() => router.reload()}>
              Refresh the page
            </Button>
          }
        />
      )}

      {/* Inventory stats */}
      {inventoryStats && !error && <InventoryStatsCards {...inventoryStats} />}

      {/* Inventory stats loading */}
      {!inventoryStats && !error && <SkeletonStats />}

      <div className="mt-8"></div>

      {/* Search bar */}
      {!error && <InventorySearchBar query={query} setQuery={setQuery} />}

      {/* Inventory list */}
      <InventoryTable
        error={error}
        inventory={inventory!}
        query={query}
        openModal={openModal}
        setQuery={setQuery}
      />

      {/* Infite scroll trigger */}
      {!error && <div ref={reloadDiv} className="-mt-12"></div>}

      {/* Modal */}
      {isOpen && data && (
        <InventorySidePanel
          closeModal={closeModal}
          data={data}
          goTo={goTo}
          page={page}
          updateTags={updateTags}
          tags={tags}
          handleChange={handleChange}
          removeTag={removeTag}
          addNewTag={addNewTag}
          loading={loading}
          deleteLoading={deleteLoading}
        />
      )}
    </>
  );
}
