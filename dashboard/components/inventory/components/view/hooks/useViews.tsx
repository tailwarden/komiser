import { FormEvent, useEffect, useState } from 'react';
import settingsService from '../../../../../services/settingsService';
import { ToastProps } from '../../../../toast/hooks/useToast';
import { InventoryFilterDataProps } from '../../../hooks/useInventory';

type useViewsProps = {
  setToast: (toast: ToastProps | undefined) => void;
};

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

type Pages = 'view' | 'excluded' | 'delete';

function useViews({ setToast }: useViewsProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [view, setView] = useState<ViewProps>(INITIAL_VIEW);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState<Pages>('view');

  useEffect(() => {
    settingsService.getViews().then(res => {
      if (res === Error) {
        console.log(res);
      } else {
        console.log(res);
      }
    });
  }, []);

  function populateView(newFilters: InventoryFilterDataProps[]) {
    setView(prev => ({ ...prev, filters: newFilters }));
  }

  function openModal(filters: InventoryFilterDataProps[]) {
    setView(INITIAL_VIEW);
    populateView(filters);
    setIsOpen(true);
  }

  function closeModal() {
    setIsOpen(false);
  }

  function handleChange(newData: { name: string }) {
    setView(prev => ({ ...prev, name: newData.name }));
  }

  function goTo(newPage: Pages) {
    setPage(newPage);
  }

  function saveView(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setLoading(true);
    const payload = view;
    const payloadJson = JSON.stringify(payload);

    settingsService.saveView(payloadJson).then(res => {
      if (res === Error) {
        setLoading(false);
        setToast({
          hasError: true,
          title: `Tags have been`,
          message: `The tags have been`
        });
      } else {
        setLoading(false);
        setToast({
          hasError: false,
          title: `Tags have been`,
          message: `The tags have been`
        });
      }
    });
  }

  return {
    isOpen,
    openModal,
    closeModal,
    view,
    handleChange,
    saveView,
    loading,
    page,
    goTo
  };
}

export default useViews;
