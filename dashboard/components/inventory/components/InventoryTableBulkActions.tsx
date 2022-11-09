import formatNumber from '../../../utils/formatNumber';
import Button from '../../button/Button';
import { InventoryStats } from '../hooks/useInventory';

type InventoryTableBulkActionsProps = {
  bulkItems: [] | string[];
  inventoryStats: InventoryStats | undefined;
  openBulkModal: (bulkItemsIds: string[]) => void;
};

function InventoryTableBulkActions({
  bulkItems,
  inventoryStats,
  openBulkModal
}: InventoryTableBulkActionsProps) {
  return (
    <>
      {bulkItems && bulkItems.length > 0 && (
        <div className="sticky flex items-center justify-between dark:border-t border-purplin-650 bottom-0 bg-gradient-to-r from-primary to-secondary dark:from-purplin-700 dark:to-purplin-700 w-full py-4 px-6 text-sm z-20">
          <p className="text-black-100">
            {bulkItems.length} {bulkItems.length > 1 ? 'items' : 'item'}{' '}
            {inventoryStats &&
              `out of ${formatNumber(inventoryStats.resources)}`}{' '}
            selected
          </p>
          <div className="flex gap-4">
            <Button size="lg" onClick={() => openBulkModal(bulkItems)}>
              Bulk manage tags
              <span className="flex items-center justify-center bg-black-900/10 dark:bg-black-900/20 text-xs py-1 px-2 rounded-lg">
                {formatNumber(bulkItems.length)}
              </span>
            </Button>
          </div>
        </div>
      )}
    </>
  );
}

export default InventoryTableBulkActions;
