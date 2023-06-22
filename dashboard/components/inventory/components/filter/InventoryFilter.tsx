import { NextRouter } from 'next/router';
import Button from '../../../button/Button';
import useFilterWizard from './hooks/useFilterWizard';
import FilterIcon from '../../../icons/FilterIcon';
import InventoryFilterDropdown from '../InventoryFilterDropdown';

type InventoryFilterProps = {
  router: NextRouter;
  setSkippedSearch: (number: number) => void;
};

function InventoryFilter({ router, setSkippedSearch }: InventoryFilterProps) {
  const { toggle, isOpen } = useFilterWizard({ router, setSkippedSearch });

  return (
    <div className="relative">
      <Button style="secondary" size="xs" onClick={toggle}>
        <FilterIcon width={20} height={20} />
        Filter by
      </Button>

      {/* Dropdown open */}
      {isOpen && (
        <InventoryFilterDropdown
          position={'right-0 top-12'}
          toggle={toggle}
          closeDropdownAfterAdd={true}
        />
      )}
    </div>
  );
}

export default InventoryFilter;
