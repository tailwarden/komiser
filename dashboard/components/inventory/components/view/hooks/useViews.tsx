import { NextRouter } from 'next/router';
import { FormEvent, useState } from 'react';
import settingsService from '../../../../../services/settingsService';
import { ToastProps } from '../../../../toast/hooks/useToast';
import {
  InventoryFilterDataProps,
  ViewProps
} from '../../../hooks/useInventory';

type useViewsProps = {
  setToast: (toast: ToastProps | undefined) => void;
  views: ViewProps[] | undefined;
  router: NextRouter;
};

const INITIAL_VIEW: ViewProps = {
  name: '',
  filters: [],
  exclude: []
};

type Pages = 'view' | 'excluded' | 'delete';

function useViews({ setToast, views, router }: useViewsProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [view, setView] = useState<ViewProps>(INITIAL_VIEW);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState<Pages>('view');

  function populateView(newFilters: InventoryFilterDataProps[]) {
    setView(prev => ({ ...prev, filters: newFilters }));
  }

  function findView(currentViews: ViewProps[]) {
    return currentViews.find(
      currentView => currentView.name === router.query.view
    );
  }

  function openModal(filters: InventoryFilterDataProps[]) {
    if (!router.query.view) {
      setView(INITIAL_VIEW);
      populateView(filters);
    } else {
      const viewToBeManaged = findView(views!);
      setView(viewToBeManaged!);
    }
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

    if (view) {
      setLoading(true);
      const payload = view;
      const payloadJson = JSON.stringify(payload);
      const { id } = view;

      if (router.query.view) {
        settingsService.updateView(id!, payloadJson).then(res => {
          if (res === Error) {
            setLoading(false);
            setToast({
              hasError: true,
              title: `${view.name} could not be created.`,
              message: `There was an error creating this custom view. Please refer to the logs and try again!`
            });
          } else {
            setLoading(false);
            setToast({
              hasError: false,
              title: `${view.name} has been created.`,
              message: `The custom view will now be accessible from the top navigation.`
            });
            closeModal();
          }
        });
      } else {
        settingsService.saveView(payloadJson).then(res => {
          if (res === Error) {
            setLoading(false);
            setToast({
              hasError: true,
              title: `${view.name} could not be created.`,
              message: `There was an error creating this custom view. Please refer to the logs and try again!`
            });
          } else {
            setLoading(false);
            setToast({
              hasError: false,
              title: `${view.name} has been created.`,
              message: `The custom view will now be accessible from the top navigation.`
            });
            closeModal();
          }
        });
      }
    }
  }

  function deleteView() {
    if (view) {
      setLoading(true);
      const { id } = view;

      settingsService.deleteView(id!).then(res => {
        if (res === Error) {
          setLoading(false);
          setToast({
            hasError: true,
            title: `${view.name} could not be deleted.`,
            message: `There was an error deleting this custom view. Please refer to the logs and try again!`
          });
        } else {
          setLoading(false);
          setToast({
            hasError: false,
            title: `${view.name} has been deleted.`,
            message: `The custom view has been successfully deleted.`
          });
          closeModal();
        }
      });
    }
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
    goTo,
    deleteView
  };
}

export default useViews;
