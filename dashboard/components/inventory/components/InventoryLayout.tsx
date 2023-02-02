import { ReactNode } from 'react';

type InventoryLayoutProps = {
  children: ReactNode;
};

function InventoryLayout({ children }: InventoryLayoutProps) {
  return (
    <>
      <nav className="fixed top-0 left-0 bottom-0 z-20 mt-[73px] flex w-[18rem] flex-col gap-6 bg-white p-8">
        All resources
      </nav>
      <main className="ml-[18rem]">{children}</main>
    </>
  );
}

export default InventoryLayout;
