import { useRouter } from 'next/router';
import formatNumber from '../../../utils/formatNumber';
import Button from '../../button/Button';
import { InventoryStats } from '../hooks/useInventory/types/useInventoryTypes';

type InventoryTableBulkActionsProps = {
  bulkItems: [] | string[];
  inventoryStats: InventoryStats | undefined;
  openBulkModal: (bulkItemsIds: string[]) => void;
  query: string;
  hideResourceFromCustomView: () => void;
  hideResourcesLoading: boolean;
};

function InventoryTableBulkActions({
  bulkItems,
  inventoryStats,
  openBulkModal,
  query,
  hideResourceFromCustomView,
  hideResourcesLoading
}: InventoryTableBulkActionsProps) {
  const router = useRouter();
  return (
    <>
      {bulkItems && bulkItems.length > 0 && (
        <div className="border-purplin-650 sticky bottom-0 flex w-full items-center justify-between bg-white py-4 px-6 text-sm font-medium shadow-[0px_-2px_4px_rgba(0,0,0,0.05)]">
          <p className="text-black-900">
            {bulkItems.length} {bulkItems.length > 1 ? 'resources' : 'resource'}{' '}
            {inventoryStats &&
              !query &&
              `out of ${formatNumber(inventoryStats.resources)}`}{' '}
            selected
          </p>
          <div className="flex gap-4">
            <Button
              size="sm"
              style="primary"
              onClick={() => openBulkModal(bulkItems)}
            >
              Manage tags
              <span className="ml-1 flex items-center justify-center rounded-lg bg-white/10 py-1 px-2 text-xs">
                {formatNumber(bulkItems.length)}
              </span>
            </Button>

            {router.query.view && (
              <Button
                size="sm"
                style="secondary"
                onClick={hideResourceFromCustomView}
                loading={hideResourcesLoading}
              >
                Hide from view
                <span className="ml-1 flex items-center justify-center rounded-lg bg-primary/10 py-1 px-2 text-xs">
                  {formatNumber(bulkItems.length)}
                </span>
              </Button>
            )}
          </div>
        </div>
      )}
    </>
  );
}

export default InventoryTableBulkActions;
