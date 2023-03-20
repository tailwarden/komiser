import { NextRouter } from 'next/router';
import { ChangeEvent } from 'react';
import formatNumber from '../../../utils/formatNumber';
import providers from '../../../utils/providerHelper';
import Checkbox from '../../checkbox/Checkbox';
import SkeletonInventory from '../../skeleton/SkeletonInventory';
import { ToastProps } from '../../toast/hooks/useToast';
import {
  InventoryItem,
  InventoryStats
} from '../hooks/useInventory/types/useInventoryTypes';
import InventorySearchBar from './InventorySearchBar';
import InventorySearchNoResults from './InventorySearchNoResults';
import InventoryTableBulkActions from './InventoryTableBulkActions';
import InventoryTableRow from './InventoryTableRow';
import InventoryTableTags from './InventoryTableTags';

type InventoryTableProps = {
  error: boolean;
  inventory: InventoryItem[] | [] | undefined;
  searchedInventory: InventoryItem[] | [] | undefined;
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
  hideResourceFromCustomView: () => void;
  hideResourcesLoading: boolean;
  setToast: (toast: ToastProps | undefined) => void;
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
  searchedLoading,
  hideResourceFromCustomView,
  hideResourcesLoading,
  setToast
}: InventoryTableProps) {
  return (
    <>
      {((!error && inventory && inventory.length > 0) ||
        (!error && searchedInventory)) && (
        <>
          <InventorySearchBar
            query={query}
            setQuery={setQuery}
            error={error}
            setToast={setToast}
          />
          <div className="rounded-lg rounded-t-none pb-6">
            <table className="w-full table-auto bg-white text-left text-sm text-gray-900">
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
                        className="min-w-[7rem] cursor-pointer py-4 pl-2 pr-6"
                      >
                        <div className="flex items-center gap-3">
                          <picture className="flex-shrink-0">
                            <img
                              src={providers.providerImg(item.provider)}
                              className="h-6 w-6 rounded-full"
                              alt={item.provider}
                            />
                          </picture>
                          <span>{item.provider}</span>
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer py-4 px-6"
                      >
                        <p className="w-12 xl:w-full">{item.service}</p>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="group relative cursor-pointer py-4 px-6"
                      >
                        <div className="peer h-full w-full"></div>
                        <p className="... w-56 truncate 2xl:w-96">
                          {item.name}
                        </p>
                        <div className="absolute left-4 top-12 z-10 hidden flex-col gap-2 rounded-lg bg-black-900 py-3 px-4 text-xs text-black-200 shadow-lg group-hover:flex">
                          {item.name}
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer py-4 px-6"
                      >
                        {item.region}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer py-4 px-6"
                      >
                        {item.account}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer whitespace-nowrap py-4 px-6 text-right"
                      >
                        ${formatNumber(item.cost)}
                      </td>
                      <td>
                        <InventoryTableTags
                          tags={item.tags}
                          setQuery={setQuery}
                          id={item.id}
                          bulkItems={bulkItems}
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
                        className="min-w-[7rem] cursor-pointer py-4 pl-2 pr-6"
                      >
                        <div className="flex items-center gap-3">
                          <picture className="flex-shrink-0">
                            <img
                              src={providers.providerImg(item.provider)}
                              className="h-6 w-6 rounded-full"
                              alt={item.provider}
                            />
                          </picture>
                          <span>{item.provider}</span>
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer py-4 px-6"
                      >
                        <p className="w-12 xl:w-full">{item.service}</p>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="group relative cursor-pointer py-4 px-6"
                      >
                        <div className="peer h-full w-full"></div>
                        <p className="... w-56 truncate 2xl:w-96">
                          {item.name}
                        </p>
                        <div className="absolute left-4 top-12 z-10 hidden flex-col gap-2 rounded-lg bg-black-900 py-3 px-4 text-xs text-black-200 shadow-lg group-hover:flex">
                          {item.name}
                        </div>
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer py-4 px-6"
                      >
                        {item.region}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer py-4 px-6"
                      >
                        {item.account}
                      </td>
                      <td
                        onClick={() => openModal(item)}
                        className="cursor-pointer whitespace-nowrap py-4 px-6 text-right"
                      >
                        ${formatNumber(item.cost)}
                      </td>
                      <td>
                        <InventoryTableTags
                          tags={item.tags}
                          setQuery={setQuery}
                          id={item.id}
                          bulkItems={bulkItems}
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
              hideResourceFromCustomView={hideResourceFromCustomView}
              hideResourcesLoading={hideResourcesLoading}
            />
          </div>
        </>
      )}
    </>
  );
}

export default InventoryTable;
