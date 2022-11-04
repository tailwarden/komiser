import Head from "next/head";
import { useRouter } from "next/router";
import { useRef } from "react";
import Button from "../components/button/Button";
import ErrorPage from "../components/error/ErrorPage";
import InventorySearchBar from "../components/inventory/components/InventorySearchBar";
import InventoryStatsCards from "../components/inventory/components/InventoryStatsCards";
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
      <div className="pb-24 rounded-lg rounded-t-none overflow-x-auto">
        <table className="table-auto text-sm text-left bg-white text-gray-900 w-full">
          {!error && (
            <thead>
              <tr className="border-b border-black-200/30">
                <th className="py-4 px-6">Cloud</th>
                <th className="py-4 px-6">Service</th>
                <th className="py-4 px-6">Name</th>
                <th className="py-4 px-6">Region</th>
                <th className="py-4 px-6">Account</th>
                <th className="py-4 px-6">Cost</th>
                <th className="py-4 px-6">Tags</th>
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
                    className="py-4 px-6 min-w-[7rem]"
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
                    <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg left-4 top-12 bg-white z-10 text-black-400 shadow-lg text-xs py-3 px-4">
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
                        <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg right-2 top-11 z-10 bg-white py-3 px-4 shadow-lg">
                          {item.tags.map((tag, index) => (
                            <div
                              key={index}
                              className="flex items-center gap-2 text-xs border-t border-black-150 pt-2 first:pt-0 first:border-none"
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
                                  className="hover:text-secondary"
                                >
                                  {tag.key}:
                                </span>
                              </div>
                              <span
                                onClick={() => setQuery(tag.value)}
                                className="font-medium hover:text-secondary"
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

        {/* Inventory loading */}
        {!inventory && !error && <SkeletonInventory />}
      </div>

      {/* Infite scroll trigger */}
      {!error && <div ref={reloadDiv} className="-mt-12"></div>}

      {/* Modal */}
      {isOpen && (
        <>
          <div
            onClick={closeModal}
            className="hidden sm:block fixed inset-0 z-30 bg-black-900/10 opacity-0 animate-fade-in"
          ></div>
          <div className="fixed overflow-auto inset-0 z-30 sm:top-4 sm:bottom-4 sm:right-4 sm:left-auto w-full sm:w-[32rem] p-6 sm:rounded-lg shadow-2xl opacity-0 animate-fade-in-up sm:animate-fade-in-left bg-white">
            {/* Modal headers */}
            <div className="flex flex-wrap-reverse sm:flex-nowrap items-center justify-between gap-6">
              {data && (
                <div className="flex flex-wrap sm:flex-nowrap items-center gap-4">
                  <picture className="flex-shrink-0">
                    <img
                      src={providers.providerImg(data.provider)}
                      className="w-8 h-8 rounded-full"
                      alt={data.provider}
                    />
                  </picture>

                  <div className="flex flex-col gap-1 py-1">
                    <p className="font-medium text-black-900 w-48 truncate ...">
                      {data.service}
                    </p>
                    <p className="text-xs text-black-300">{data.name}</p>
                  </div>
                </div>
              )}

              <div className="flex gap-4 flex-shrink-0">
                <Button style="secondary" onClick={closeModal}>
                  Close
                </Button>
              </div>
            </div>

            {/* Tabs */}
            <div className="mt-4"></div>
            <div className="text-sm font-medium text-center border-b-2 border-black-150 text-black-300">
              <ul className="flex justify-between sm:justify-start -mb-[2px]">
                <li className="mr-2">
                  <a
                    onClick={() => goTo("tags")}
                    className={`select-none inline-block py-4 px-2 sm:p-4 rounded-t-lg border-b-2 border-transparent hover:text-komiser-700 cursor-pointer 
                        ${
                          (page === "tags" || page === "delete") &&
                          `text-secondary border-secondary`
                        }`}
                  >
                    Tags
                  </a>
                </li>
              </ul>
            </div>

            {/* Tags form */}
            <div className="mt-6"></div>
            <div className="p-6 bg-black-100 rounded-lg">
              <div className="flex flex-col gap-6">
                {page === "tags" && (
                  <form
                    onSubmit={(e) => {
                      e.preventDefault();

                      updateTags();
                    }}
                    className="flex flex-col gap-6"
                  >
                    {tags &&
                      tags.map((tag, id) => (
                        <div key={id} className="flex gap-6">
                          <TagWrapper
                            tag={tag}
                            id={id}
                            handleChange={handleChange}
                          />
                          {tags.length > 1 && (
                            <Button
                              size="xs"
                              style="ghost"
                              onClick={() => removeTag(id)}
                            >
                              <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="20"
                                height="20"
                                fill="none"
                                viewBox="0 0 24 24"
                              >
                                <path
                                  stroke="currentColor"
                                  strokeLinecap="round"
                                  strokeLinejoin="round"
                                  strokeWidth="2"
                                  d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
                                ></path>
                              </svg>
                            </Button>
                          )}
                        </div>
                      ))}
                    <div
                      onClick={addNewTag}
                      className="flex items-center justify-center gap-2 py-3 bg-white hover:bg-komiser-700/10 rounded-lg text-black-900/50 text-sm transition-colors cursor-pointer"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="none"
                        viewBox="0 0 24 24"
                      >
                        <path
                          fill="currentColor"
                          d="M18 12.75H6c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h12c.41 0 .75.34.75.75s-.34.75-.75.75z"
                        ></path>
                        <path
                          fill="currentColor"
                          d="M12 18.75c-.41 0-.75-.34-.75-.75V6c0-.41.34-.75.75-.75s.75.34.75.75v12c0 .41-.34.75-.75.75z"
                        ></path>
                      </svg>
                      Add new tag
                    </div>
                    <div className="flex items-center justify-between">
                      <Button
                        type="submit"
                        size="lg"
                        loading={loading}
                        disabled={
                          tags &&
                          !tags.every(
                            (tag) => tag.key.trim() && tag.value.trim()
                          )
                        }
                      >
                        {data && data.tags && data.tags.length > 0
                          ? "Save changes"
                          : "Add tags"}
                      </Button>
                      {data && data.tags && data.tags.length > 0 && (
                        <Button
                          size="lg"
                          style="delete"
                          loading={deleteLoading}
                          onClick={() => {
                            updateTags("delete");
                          }}
                        >
                          Delete all tags
                        </Button>
                      )}
                    </div>
                  </form>
                )}
              </div>
            </div>
          </div>
        </>
      )}
    </>
  );
}
