import { NextRouter } from 'next/router';
import { ViewProps } from '../hooks/useInventory';

type InventoryViewsHeaderProps = {
  views: ViewProps[] | undefined;
  router: NextRouter;
};

function InventoryViewsHeader({ views, router }: InventoryViewsHeaderProps) {
  const currentView = views?.find(
    view => view.id.toString() === router.query.view
  );
  return (
    <>
      <p className="text-lg font-medium text-black-900">
        {currentView ? currentView.name : 'All resources'}
      </p>
    </>
  );
}

export default InventoryViewsHeader;
