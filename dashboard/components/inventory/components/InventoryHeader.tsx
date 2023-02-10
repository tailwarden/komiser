import { ReactNode } from 'react';

type InventoryHeaderProps = {
  isNotCustomView: boolean;
  children: ReactNode;
};

function InventoryHeader({ isNotCustomView, children }: InventoryHeaderProps) {
  return (
    <div className="flex min-h-[40px] items-center justify-between gap-8">
      {isNotCustomView && (
        <p className="flex items-center gap-2 text-lg font-medium text-black-900">
          All Resources
        </p>
      )}
      <div className="flex flex-shrink-0 items-center gap-4">{children}</div>
    </div>
  );
}

export default InventoryHeader;
