import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { ChangeEvent, FormEvent, useState } from 'react';
import settingsService from '../../../../../services/settingsService';
import {
  HiddenResource,
  InventoryFilterData,
  View
} from '../../../hooks/useInventory/types/useInventoryTypes';

type useViewsProps = {
  showToast: (toast: ToastProps) => void;
  views: View[] | undefined;
  router: NextRouter;
  getViews: (
    edit?: boolean | undefined,
    viewName?: string | undefined,
    redirect?: boolean | undefined
  ) => void;
  hiddenResources: HiddenResource[] | undefined;
  setHideOrUnhideHasUpdate: (hideOrUnhideHasUpdate: boolean) => void;
};

const INITIAL_VIEW: View = {
  id: 0,
  name: '',
  filters: [],
  exclude: []
};

export type ViewsPages =
  | 'view'
  | 'excluded'
  | 'delete'
  | 'hidden resources'
  | 'alerts';

function useViews({
  showToast,
  views,
  router,
  getViews,
  hiddenResources,
  setHideOrUnhideHasUpdate
}: useViewsProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [view, setView] = useState<View>(INITIAL_VIEW);
  const [loading, setLoading] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [page, setPage] = useState<ViewsPages>('view');
  const [bulkItems, setBulkItems] = useState<number[] | []>([]);
  const [bulkSelectCheckbox, setBulkSelectCheckbox] = useState(false);
  const [unhideLoading, setUnhideLoading] = useState(false);

  function findView(currentViews: View[]) {
    return currentViews.find(
      currentView => currentView.id.toString() === router.query.view
    );
  }

  function populateView(newFilters: InventoryFilterData[]) {
    setView(prev => ({ ...prev, filters: newFilters }));
  }

  function openModal(filters?: InventoryFilterData[], openPage?: ViewsPages) {
    setPage('view');
    setBulkItems([]);
    setBulkSelectCheckbox(false);

    if (!router.query.view && filters) {
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

  function goTo(newPage: ViewsPages) {
    setPage(newPage);
  }

  function saveView(
    e: FormEvent<HTMLFormElement>,
    duplicate?: boolean,
    viewToBeDuplicated?: View
  ) {
    e.preventDefault();

    if (view && !duplicate) {
      setLoading(true);
      const payload = view;
      const payloadJson = JSON.stringify(payload);
      const { id } = view;

      if (router.query.view) {
        settingsService.updateView(id.toString(), payloadJson).then(res => {
          if (res === Error) {
            setLoading(false);
            showToast({
              hasError: true,
              title: `${view.name} could not be saved.`,
              message: `There was an error saving this custom view. Please refer to the logs and try again!`
            });
          } else {
            setLoading(false);
            getViews(true, view.id.toString());
            showToast({
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
            showToast({
              hasError: true,
              title: `${view.name} could not be created.`,
              message: `There was an error creating this custom view. Please refer to the logs and try again!`
            });
          } else {
            setLoading(false);
            getViews(false, undefined, true);
            showToast({
              hasError: false,
              title: `${view.name} has been created.`,
              message: `The custom view will now be accessible from the side navigation.`
            });
            closeModal();
          }
        });
      }
    }

    if (duplicate && viewToBeDuplicated) {
      setLoading(true);
      const payload = viewToBeDuplicated;
      payload.id = 0;
      payload.name = `(copy) ${payload.name}`;
      const payloadJson = JSON.stringify(payload);
      settingsService.saveView(payloadJson).then(res => {
        if (res === Error) {
          setLoading(false);
          showToast({
            hasError: true,
            title: `${viewToBeDuplicated.name} could not be duplicated.`,
            message: `There was an error duplicating this custom view. Please refer to the logs and try again!`
          });
        } else {
          setLoading(false);
          getViews(false, undefined, true);
          showToast({
            hasError: false,
            title: `${viewToBeDuplicated.name} has been duplicated.`,
            message: `The custom view will now be accessible from the side navigation.`
          });
          closeModal();
        }
      });
    }
  }

  function deleteView(dropdown?: boolean, viewToBeDeleted?: View) {
    if (view && !dropdown) {
      setLoading(true);
      const { id } = view;

      settingsService.deleteView(id.toString()).then(res => {
        if (res === Error) {
          setLoading(false);
          showToast({
            hasError: true,
            title: `${view.name} could not be deleted.`,
            message: `There was an error deleting this custom view. Please refer to the logs and try again!`
          });
        } else {
          getViews();
          setLoading(false);
          showToast({
            hasError: false,
            title: `${view.name} has been deleted.`,
            message: `The custom view has been successfully deleted.`
          });
          closeModal();
          router.push(router.pathname);
        }
      });
    }

    if (dropdown && viewToBeDeleted) {
      setDeleteLoading(true);
      const { id } = viewToBeDeleted;

      settingsService.deleteView(id.toString()).then(res => {
        if (res === Error) {
          setDeleteLoading(false);
          showToast({
            hasError: true,
            title: `${viewToBeDeleted.name} could not be deleted.`,
            message: `There was an error deleting this custom view. Please refer to the logs and try again!`
          });
        } else {
          getViews();
          setDeleteLoading(false);
          showToast({
            hasError: false,
            title: `${viewToBeDeleted.name} has been deleted.`,
            message: `The custom view has been successfully deleted.`
          });
          router.push(router.pathname);
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
        showToast({
          hasError: true,
          title: 'Resources could not be unhid.',
          message:
            'There was an error unhiding the resources. Please refer to the logs and try again.'
        });
      } else {
        setUnhideLoading(false);
        showToast({
          hasError: false,
          title: 'Resources are now unhidden.',
          message: 'The resources were successfully unhidden.'
        });
        setHideOrUnhideHasUpdate(true);
        closeModal();
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
    unhideResources,
    deleteLoading
  };
}

export default useViews;
