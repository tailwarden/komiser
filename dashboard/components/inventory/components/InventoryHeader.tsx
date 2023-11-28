import { ReactNode } from 'react';

type InventoryHeaderProps = {
  isNotCustomView: boolean;
};

function InventoryHeader({ isNotCustomView }: InventoryHeaderProps) {
  return (
    <div className="flex min-h-[40px] items-center justify-between gap-8">
      {isNotCustomView && (
        <p className="flex items-center gap-2 text-lg font-medium text-gray-950">
          All Resources
        </p>
      )}
    </div>
  );
}

export default InventoryHeader;
