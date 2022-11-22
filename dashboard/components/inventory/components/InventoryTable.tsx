import { NextRouter } from 'next/router';
import React, { ChangeEvent } from 'react';
import formatNumber from '../../../utils/formatNumber';
import providers from '../../../utils/providerHelper';
import Checkbox from '../../checkbox/Checkbox';
import SkeletonInventory from '../../skeleton/SkeletonInventory';
import { InventoryItem, InventoryStats } from '../hooks/useInventory';
import InventorySearchBar from './InventorySearchBar';
import InventorySearchNoResults from './InventorySearchNoResults';
import InventoryTableBulkActions from './InventoryTableBulkActions';
import InventoryTableRow from './InventoryTableRow';
import InventoryTableTags from './InventoryTableTags';

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
  inventoryStats: InventoryStats | undefined;
  openBulkModal: (bulkItemsIds: string[]) => void;
  router: NextRouter;
  searchedLoading: boolean;
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
  onCheckboxChange,
  inventoryStats,
  openBulkModal,
  router,
  searchedLoading
}: InventoryTableProps) {
  return (
    <>
      {((!error && inventory && inventory.length > 0) ||
        (!error && searchedInventory)) && (
        <>
          <InventorySearchBar query={query} setQuery={setQuery} error={error} />
          <div className="pb-6 rounded-lg rounded-t-none">
            <table className="table-auto text-sm text-left bg-white text-gray-900 w-full">
              {!error && (
                <thead className="sticky top-[73px] z-10 bg-white">
                  <tr className="shadow-[inset_0_-1px_0_0_#cfd7d74d]">
                    <th className="py-4 pl-6">
                      <div className="flex items-center">
                        <Checkbox
                          checked={bulkSelectCheckbox}
                          onChange={handleBulkSelection}
                        />
                      </div>
                    </th>
                    <th className="pl-2 pr-6">Cloud</th>
                    <th className="py-4 px-6">Service</th>
                    <th className="py-4 px-6">Name</th>
                    <th className="py-4 px-6">Region</th>
                    <th className="py-4 px-6">Account</th>
                    <th className="py-4 px-6 text-right">Cost</th>
                    <th className="py-4 px-6">Tags</th>
                  </tr>
                </thead>
              )}
              <tbody>
                {/* Inventory table */}
                {!query &&
                  !searchedInventory &&
                  inventory &&
                  inventory.length > 0 &&
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
                        className="py-4 pl-2 pr-6 min-w-[7rem] cursor-pointer"
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
                        <p className="w-56 2xl:w-96 truncate ...">
                          {item.name}
                        </p>
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
                        className="py-4 px-6 whitespace-nowrap cursor-pointer text-right"
                      >
                        ${formatNumber(item.cost)}
                      </td>
                      <td>
                        <InventoryTableTags
                          tags={item.tags}
                          setQuery={setQuery}
                        />
                      </td>
                    </InventoryTableRow>
                  ))}

                {/* Searched inventory table */}
                {!searchedLoading &&
                  searchedInventory &&
                  searchedInventory.length > 0 &&
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
                        className="py-4 pl-2 pr-6 min-w-[7rem] cursor-pointer"
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
                        <p className="w-56 xl:w-72 2xl:w-96 truncate ...">
                          {item.name}
                        </p>
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
                        className="py-4 px-6 whitespace-nowrap cursor-pointer text-right"
                      >
                        ${formatNumber(item.cost)}
                      </td>
                      <td>
                        <InventoryTableTags
                          tags={item.tags}
                          setQuery={setQuery}
                        />
                      </td>
                    </InventoryTableRow>
                  ))}
              </tbody>
            </table>

            {/* Inventory search loading */}
            {searchedLoading && <SkeletonInventory />}

            {/* Inventory search no results */}
            {searchedInventory &&
              searchedInventory.length === 0 &&
              !searchedLoading && (
                <InventorySearchNoResults
                  query={query}
                  setQuery={setQuery}
                  router={router}
                />
              )}

            {/* Bulk actions sticky footer */}
            <InventoryTableBulkActions
              bulkItems={bulkItems}
              inventoryStats={inventoryStats}
              openBulkModal={openBulkModal}
              query={query}
            />
          </div>
        </>
      )}
    </>
  );
}

export default InventoryTable;
