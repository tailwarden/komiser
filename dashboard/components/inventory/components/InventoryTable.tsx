import React, { ChangeEvent } from 'react';
import formatNumber from '../../../utils/formatNumber';
import providers from '../../../utils/providerHelper';
import Checkbox from '../../checkbox/Checkbox';
import SkeletonInventory from '../../skeleton/SkeletonInventory';
import { InventoryItem } from '../hooks/useInventory';
import InventorySearchBar from './InventorySearchBar';
import InventorySearchNoResults from './InventorySearchNoResults';
import InventoryTableRow from './InventoryTableRow';

type InventoryTableProps = {
  error: boolean;
  inventory: InventoryItem[] | [];
  searchedInventory: InventoryItem[] | [];
  query: string;
  openModal: (item: InventoryItem) => void;
  setQuery: (query: string) => void;
  bulkSelectCheckbox: boolean;
  handleBulkSelection: (e: ChangeEvent<HTMLInputElement>) => void;
  bulkItems: [] | string[];
  onCheckboxChange: (e: ChangeEvent<HTMLInputElement>, id: string) => void;
};

function InventoryTable({
  error,
  inventory,
  searchedInventory,
  query,
  openModal,
  setQuery,
  bulkSelectCheckbox,
  handleBulkSelection,
  bulkItems,
  onCheckboxChange
}: InventoryTableProps) {
  return (
    <>
      {inventory && Object.keys(inventory).length !== 0 && !error && (
        <>
          <InventorySearchBar query={query} setQuery={setQuery} error={error} />
          <div className="pb-24 rounded-lg rounded-t-none">
            <table className="table-auto text-sm text-left bg-white text-gray-900 w-full">
              {!error && (
                <thead>
                  <tr className="border-b border-black-200/30">
                    <th className="py-4 pl-6">
                      <div className="flex items-center">
                        <Checkbox
                          checked={bulkSelectCheckbox}
                          onChange={handleBulkSelection}
                        />
                      </div>
                    </th>
                    <th className="pl-4 pr-6">Cloud</th>
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
                {!query &&
                  inventory.map(item => (
                    <InventoryTableRow
                      key={item.id}
                      id={item.id}
                      bulkItems={bulkItems}
                    >
                      <td className="py-4 pl-6">
                        <div className="flex items-center">
                          <Checkbox
                            checked={
                              bulkItems &&
                              !!bulkItems.find(
                                currentId => currentId === item.id
                              )
                            }
                            onChange={e => onCheckboxChange(e, item.id)}
                          />
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 pl-4 pr-6 min-w-[7rem] cursor-pointer"
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
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 cursor-pointer"
                      >
                        <p className="w-12 xl:w-full">{item.service}</p>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 group relative cursor-pointer"
                      >
                        <div className="peer w-full h-full"></div>
                        <p className="w-48 truncate ...">{item.name}</p>
                        <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg left-4 top-12 bg-black-900 z-10 text-black-200 shadow-lg text-xs py-3 px-4">
                          {item.name}
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 cursor-pointer"
                      >
                        {item.region}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 cursor-pointer"
                      >
                        {item.account}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 whitespace-nowrap cursor-pointer"
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
                            <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg right-6 top-11 z-10 bg-white py-3 px-4 shadow-lg">
                              {item.tags.map((tag, index) => (
                                <div
                                  key={index}
                                  className="flex items-center gap-2 text-xs border-t border-black-150 -mx-4 px-4 pt-2 first:pt-0 first:border-none"
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
                                      onClick={e => {
                                        setQuery(tag.key);
                                      }}
                                      className="hover:text-secondary cursor-pointer"
                                    >
                                      {tag.key}:
                                    </span>
                                  </div>
                                  <span
                                    onClick={() => setQuery(tag.value)}
                                    className="font-medium hover:text-secondary cursor-pointer"
                                  >
                                    {tag.value}
                                  </span>
                                </div>
                              ))}
                            </div>
                          </div>
                        )}
                      </td>
                    </InventoryTableRow>
                  ))}

                {query &&
                  searchedInventory &&
                  Object.keys(searchedInventory).length !== 0 &&
                  searchedInventory.map(item => (
                    <InventoryTableRow
                      key={item.id}
                      id={item.id}
                      bulkItems={bulkItems}
                    >
                      <td className="py-4 pl-6">
                        <div className="flex items-center">
                          <Checkbox
                            checked={
                              bulkItems &&
                              !!bulkItems.find(
                                currentId => currentId === item.id
                              )
                            }
                            onChange={e => onCheckboxChange(e, item.id)}
                          />
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 pl-4 pr-6 min-w-[7rem] cursor-pointer"
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
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 cursor-pointer"
                      >
                        <p className="w-12 xl:w-full">{item.service}</p>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 group relative cursor-pointer"
                      >
                        <div className="peer w-full h-full"></div>
                        <p className="w-48 truncate ...">{item.name}</p>
                        <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg left-4 top-12 bg-black-900 z-10 text-black-200 shadow-lg text-xs py-3 px-4">
                          {item.name}
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 cursor-pointer"
                      >
                        {item.region}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 cursor-pointer"
                      >
                        {item.account}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="py-4 px-6 whitespace-nowrap cursor-pointer"
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
                            <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg right-6 top-11 z-10 bg-white py-3 px-4 shadow-lg">
                              {item.tags.map((tag, index) => (
                                <div
                                  key={index}
                                  className="flex items-center gap-2 text-xs border-t border-black-150 -mx-4 px-4 pt-2 first:pt-0 first:border-none"
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
                                      onClick={e => {
                                        setQuery(tag.key);
                                      }}
                                      className="hover:text-secondary cursor-pointer"
                                    >
                                      {tag.key}:
                                    </span>
                                  </div>
                                  <span
                                    onClick={() => setQuery(tag.value)}
                                    className="font-medium hover:text-secondary cursor-pointer"
                                  >
                                    {tag.value}
                                  </span>
                                </div>
                              ))}
                            </div>
                          </div>
                        )}
                      </td>
                    </InventoryTableRow>
                  ))}
              </tbody>
            </table>

            {/* Inventory search loading */}
            {query && !searchedInventory && <SkeletonInventory />}

            {/* Inventory search no results */}
            {query && searchedInventory && searchedInventory.length === 0 && (
              <InventorySearchNoResults query={query} setQuery={setQuery} />
            )}
          </div>
        </>
      )}
    </>
  );
}

export default InventoryTable;
