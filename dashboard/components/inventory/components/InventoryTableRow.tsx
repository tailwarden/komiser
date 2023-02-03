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
          ? 'border-black-200/70 bg-komiser-120'
          : 'border-black-200/30 bg-white hover:bg-black-100/50'
      } border-b last:border-none`}
    >
      {children}
    </tr>
  );
}

export default InventoryTableRow;
