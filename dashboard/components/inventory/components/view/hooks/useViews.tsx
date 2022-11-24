import { useState } from 'react';
import { InventoryFilterDataProps } from '../../../hooks/useInventory';

type ViewProps = {
  name: string;
  filters: InventoryFilterDataProps[];
  exclude: string[];
};

const INITIAL_VIEW: ViewProps = {
  name: '',
  filters: [],
  exclude: []
};

function useViews() {
  const [isOpen, setIsOpen] = useState(false);
  const [view, setView] = useState<ViewProps>(INITIAL_VIEW);

  function populateView(newFilters: InventoryFilterDataProps[]) {
    setView(prev => ({ ...prev, filters: newFilters }));
  }

  function openModal(filters: InventoryFilterDataProps[]) {
    populateView(filters);
    setIsOpen(true);
  }

  function closeModal() {
    setIsOpen(false);
  }

  function handleChange(newData: { name: string }) {
    setView(prev => ({ ...prev, name: newData.name }));
  }

  return {
    isOpen,
    openModal,
    closeModal,
    view,
    handleChange
  };
}

export default useViews;
