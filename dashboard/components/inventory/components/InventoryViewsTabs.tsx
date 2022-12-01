import { NextRouter } from 'next/router';
import { ViewProps } from '../hooks/useInventory';

type InventoryViewsTabsProps = {
  views: ViewProps[] | undefined;
  router: NextRouter;
};

function InventoryViewsTabs({ views, router }: InventoryViewsTabsProps) {
  return (
    <>
      <div className="text-sm font-medium text-center text-black-300">
        <ul className="flex flex-wrap justify-between sm:justify-start -mb-[2px]">
          <li>
            <a
              onClick={() => router.push('/')}
              className={`select-none inline-block py-4 px-2 sm:p-4 rounded-t-lg border-b-2 border-transparent hover:text-komiser-700 cursor-pointer 
                       ${
                         !router.query.view &&
                         `text-komiser-600 border-komiser-600 hover:text-komiser-600`
                       }`}
            >
              Inventory
            </a>
          </li>
          {views &&
            views.length > 0 &&
            views.map((view, idx) => (
              <li key={idx}>
                <a
                  onClick={() => router.push(`/?view=${view.name}`)}
                  className={`select-none inline-block py-4 px-2 sm:p-4 rounded-t-lg border-b-2 border-transparent hover:text-komiser-700 cursor-pointer whitespace-nowrap
                       ${
                         router.asPath.replaceAll('%20', ' ') ===
                           `/?view=${view.name}` &&
                         `text-komiser-600 border-komiser-600 hover:text-komiser-600`
                       }`}
                >
                  {view.name}
                </a>
              </li>
            ))}
        </ul>
      </div>
    </>
  );
}

export default InventoryViewsTabs;
