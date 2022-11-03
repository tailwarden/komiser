import Head from "next/head";
import { useRef } from "react";
import ErrorPage from "../components/error/ErrorPage";
import InventoryStatsCards from "../components/inventory/components/InventoryStatsCards";
import useInventory from "../components/inventory/hooks/useInventory";
import SkeletonStats from "../components/skeleton/SkeletonStats";
import Toast from "../components/toast/Toast";
import formatNumber from "../utils/formatNumber";
import providers from "../utils/providerHelper";

export default function Inventory() {
  const inventoryStats = {
    resources: 140,
    cost: 33242,
    savings: 55,
    regions: 8,
  };
  const reloadDiv = useRef<HTMLDivElement>(null);
  const {
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

      {/* Error page */}
      {((error && !inventoryStats) || (error && !inventory)) && (
        <ErrorPage
          title="Network request error."
          message="There was an error fetching the inventory data. Please refresh the page."
        />
      )}

      {/* Inventory stats */}
      {inventoryStats && !error && <InventoryStatsCards {...inventoryStats} />}

      {/* Inventory stats loading */}
      {!inventoryStats && !error && <SkeletonStats />}

      <div className="mt-8"></div>
      <div className="pb-24 rounded-lg rounded-t-none">
        <table className="table-auto text-sm text-left bg-white text-gray-900 w-full">
          {!error && (
            <thead>
              <tr className="border-b border-black-200/30">
                <th className="py-4 pl-4 pr-6">Cloud</th>
                <th className="py-4 px-6">Service</th>
                <th className="py-4 px-6">Name</th>
                <th className="py-4 px-6">Region</th>
                <th className="py-4 px-6">Account</th>
                <th className="py-4 px-6">Cost</th>
                <th className="py-4 px-6"></th>
              </tr>
            </thead>
          )}
          <tbody>
            {/* Inventory table */}
            {inventory &&
              !error &&
              !query &&
              inventory.map((item) => (
                <tr
                  key={item.id}
                  className="bg-white hover:bg-black-100/50 border-b border-black-200/30 last:border-none cursor-pointer"
                >
                  <td
                    onClick={() => openModal(item)}
                    className="py-4 pl-4 pr-6 min-w-[7rem]"
                  >
                    <div className="flex items-center gap-3">
                      <picture className="flex-shrink-0">
                        <img
                          src={providers.providerImg(item.provider)}
                          className="w-6 h-6 rounded-full"
                          alt={item.provider}
                        />
                      </picture>
                      <span>{item.provider}</span>
                    </div>
                  </td>
                  <td onClick={() => openModal(item)} className="py-4 px-6">
                    <p className="w-12 xl:w-full">{item.service}</p>
                  </td>
                  <td
                    onClick={() => openModal(item)}
                    className="py-4 px-6 group relative"
                  >
                    <div className="peer w-full h-full"></div>
                    <p className="w-48 truncate ...">{item.name}</p>
                    <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg left-4 top-12 bg-white z-10 dark:bg-purplin-900 text-black-400 shadow-lg dark:shadow-none text-xs py-3 px-4">
                      {item.name}
                    </div>
                  </td>
                  <td onClick={() => openModal(item)} className="py-4 px-6">
                    {item.region}
                  </td>
                  <td onClick={() => openModal(item)} className="py-4 px-6">
                    {item.account}
                  </td>
                  <td
                    onClick={() => openModal(item)}
                    className="py-4 px-6 whitespace-nowrap"
                  >
                    ${formatNumber(item.cost)}
                  </td>
                  <td>
                    {item.tags && item.tags.length > 0 && (
                      <div className="relative group">
                        <div className="flex items-center gap-1 py-4 px-6">
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="16"
                            height="16"
                            fill="none"
                            viewBox="0 0 24 24"
                          >
                            <path
                              stroke="currentColor"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth="2"
                              d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
                            ></path>
                            <path
                              stroke="currentColor"
                              strokeLinecap="round"
                              strokeWidth="2"
                              d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
                            ></path>
                          </svg>
                          <span>{item.tags.length}</span>
                        </div>
                        <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg right-2 top-11 z-10 bg-white dark:bg-purplin-900 py-3 px-4 shadow-lg dark:shadow-none">
                          {item.tags.map((tag, index) => (
                            <div
                              key={index}
                              className="flex items-center gap-2 text-xs border-t border-black-150  dark:border-black-400/50 pt-2 first:pt-0 first:border-none"
                            >
                              <div className="flex items-center gap-1 text-black-300">
                                <svg
                                  xmlns="http://www.w3.org/2000/svg"
                                  width="14"
                                  height="14"
                                  fill="none"
                                  viewBox="0 0 24 24"
                                >
                                  <path
                                    stroke="currentColor"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth="2"
                                    d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
                                  ></path>
                                  <path
                                    stroke="currentColor"
                                    strokeLinecap="round"
                                    strokeWidth="2"
                                    d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
                                  ></path>
                                </svg>
                                <span
                                  onClick={(e) => {
                                    setQuery(tag.key);
                                  }}
                                  className="hover:text-primary dark:hover:text-purplin-300"
                                >
                                  {tag.key}:
                                </span>
                              </div>
                              <span
                                onClick={() => setQuery(tag.value)}
                                className="font-medium hover:text-primary dark:hover:text-purplin-300"
                              >
                                {tag.value}
                              </span>
                            </div>
                          ))}
                        </div>
                      </div>
                    )}
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </div>
    </>
  );
}
