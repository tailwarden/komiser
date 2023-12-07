import { ReactNode } from 'react';

type InventoryTableRowProps = {
  id: string;
  children: ReactNode;
  bulkItems: [] | string[];
};
function InventoryTableRow({
  id,
  children,
  bulkItems
}: InventoryTableRowProps) {
  return (
    <tr
      className={`${
        bulkItems && bulkItems.find(currentId => currentId === id)
          ? 'border-gray-300 bg-darkcyan-100'
          : 'border-gray-300 bg-white hover:bg-gray-50'
      } border-b last:border-none`}
    >
      {children}
    </tr>
  );
}

export default InventoryTableRow;
