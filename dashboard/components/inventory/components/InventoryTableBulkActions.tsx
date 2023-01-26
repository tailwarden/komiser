import { useRouter } from 'next/router';
import formatNumber from '../../../utils/formatNumber';
import Button from '../../button/Button';
import { InventoryStats } from '../hooks/useInventory';

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
        <div className="sticky flex items-center justify-between border-purplin-650 bottom-0 bg-gradient-to-r from-komiser-600 to-komiser-700 w-full py-4 px-6 text-sm">
          <p className="text-black-100">
            {bulkItems.length} {bulkItems.length > 1 ? 'resources' : 'resource'}{' '}
            {inventoryStats &&
              !query &&
              `out of ${formatNumber(inventoryStats.resources)}`}{' '}
            selected
          </p>
          <div className="flex gap-4">
            <Button
              size="lg"
              style="bulk"
              onClick={() => openBulkModal(bulkItems)}
            >
              Manage tags
              <span className="flex items-center justify-center bg-primary/10 text-xs py-1 px-2 rounded-lg">
                {formatNumber(bulkItems.length)}
              </span>
            </Button>

            {router.query.view && (
              <Button
                size="lg"
                style="bulk-outline"
                onClick={hideResourceFromCustomView}
                loading={hideResourcesLoading}
              >
                Hide from view
                <span className="flex items-center justify-center bg-white/10 text-xs py-1 px-2 rounded-lg">
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
