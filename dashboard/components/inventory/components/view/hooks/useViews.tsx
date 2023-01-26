import { NextRouter } from 'next/router';
import { ChangeEvent, FormEvent, useState } from 'react';
import settingsService from '../../../../../services/settingsService';
import { ToastProps } from '../../../../toast/hooks/useToast';
import {
  HiddenResource,
  InventoryFilterDataProps,
  ViewProps
} from '../../../hooks/useInventory';

type useViewsProps = {
  setToast: (toast: ToastProps | undefined) => void;
  views: ViewProps[] | undefined;
  router: NextRouter;
  getViews: (edit?: boolean | undefined, viewName?: string | undefined) => void;
  hiddenResources: HiddenResource[] | undefined;
};

const INITIAL_VIEW: ViewProps = {
  id: 0,
  name: '',
  filters: [],
  exclude: []
};

type Pages = 'view' | 'excluded' | 'delete' | 'hidden resources';

function useViews({
  setToast,
  views,
  router,
  getViews,
  hiddenResources
}: useViewsProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [view, setView] = useState<ViewProps>(INITIAL_VIEW);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState<Pages>('view');
  const [bulkItems, setBulkItems] = useState<number[] | []>([]);
  const [bulkSelectCheckbox, setBulkSelectCheckbox] = useState(false);
  const [unhideLoading, setUnhideLoading] = useState(false);

  function findView(currentViews: ViewProps[]) {
    return currentViews.find(
      currentView => currentView.id.toString() === router.query.view
    );
  }

  function populateView(newFilters: InventoryFilterDataProps[]) {
    setView(prev => ({ ...prev, filters: newFilters }));
  }

  function openModal(filters: InventoryFilterDataProps[], openPage?: Pages) {
    setPage('view');
    setBulkItems([]);
    setBulkSelectCheckbox(false);

    if (!router.query.view) {
      setView(INITIAL_VIEW);
      populateView(filters);
    } else {
      const viewToBeManaged = findView(views!);
      setView(viewToBeManaged!);
    }

    if (openPage) {
      setPage(openPage);
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
              title: `${view.name} could not be saved.`,
              message: `There was an error saving this custom view. Please refer to the logs and try again!`
            });
          } else {
            setLoading(false);
            getViews(true, view.id.toString());
            setToast({
              hasError: false,
              title: `${view.name} has been saved.`,
              message: `The custom view has been successfully saved.`
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
            getViews();
            setToast({
              hasError: false,
              title: `${view.name} has been created.`,
              message: `The custom view will now be accessible from the top navigation.`
            });
            closeModal();
            router.push('/');
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
          getViews();
          setLoading(false);
          setToast({
            hasError: false,
            title: `${view.name} has been deleted.`,
            message: `The custom view has been successfully deleted.`
          });
          closeModal();
          router.push('/');
        }
      });
    }
  }

  function onCheckboxChange(e: ChangeEvent<HTMLInputElement>, id: number) {
    if (e.target.checked) {
      const newArray = [...bulkItems];
      newArray.push(id);
      setBulkItems(newArray);
    } else {
      const newArray = bulkItems.filter(currentId => currentId !== id);
      setBulkItems(newArray);
    }
  }

  function handleBulkSelection(e: ChangeEvent<HTMLInputElement>) {
    if (e.target.checked && hiddenResources) {
      const arrayOfIds = hiddenResources.map(item => item.id);
      setBulkItems(arrayOfIds);
      setBulkSelectCheckbox(true);
    } else {
      setBulkItems([]);
      setBulkSelectCheckbox(false);
    }
  }

  function unhideResources() {
    if (!router.query.view || bulkItems.length === 0) return;

    const hiddenResourcesIds: number[] = hiddenResources!.map(
      resource => resource.id
    );
    const checkboxIds: number[] = bulkItems;

    const idsToExclude = hiddenResourcesIds!.filter(
      id => checkboxIds.indexOf(id) === -1
    );

    const viewId = router.query.view.toString();
    const newPayload = { id: Number(viewId), exclude: idsToExclude };
    const payload = JSON.stringify(newPayload);

    settingsService.unhideResourceFromView(viewId, payload).then(res => {
      if (res === Error) {
        setUnhideLoading(false);
        setToast({
          hasError: true,
          title: 'Resources could not be hided.',
          message:
            'There was an error hiding the resources. Please refer to the logs and try again.'
        });
      } else {
        setUnhideLoading(false);
        setToast({
          hasError: false,
          title: 'Resources are now hidden.',
          message:
            'The resources were successfully hidden. They can be unhided from the custom view management.'
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
    goTo,
    deleteView,
    bulkItems,
    bulkSelectCheckbox,
    onCheckboxChange,
    handleBulkSelection,
    unhideLoading,
    unhideResources
  };
}

export default useViews;
