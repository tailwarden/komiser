import { useRouter } from 'next/router';
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
  openBulkModal,
  hideResourceFromCustomView,
  hideResourcesLoading
}: InventoryTableBulkActionsProps) {
  const router = useRouter();
  const resourceText = bulkItems.length > 1 ? 'resources' : 'resource';
  return (
    <>
      {bulkItems && bulkItems.length > 0 && (
        <div className="border-purplin-650 sticky bottom-0 flex w-full items-center justify-between bg-white px-6 py-4 text-sm font-medium shadow-[0px_-2px_4px_rgba(0,0,0,0.05)]">
          <p className="text-gray-950">
            {bulkItems.length} {resourceText} {''}
            selected
          </p>
          <div className="flex gap-4">
            <Button
              size="sm"
              style="primary"
              onClick={() => openBulkModal(bulkItems)}
            >
              Tag {resourceText}
            </Button>
            {router.query.view && (
              <Button
                size="sm"
                style="secondary"
                onClick={hideResourceFromCustomView}
                loading={hideResourcesLoading}
              >
                Hide from view
              </Button>
            )}
          </div>
        </div>
      )}
    </>
  );
}

export default InventoryTableBulkActions;
