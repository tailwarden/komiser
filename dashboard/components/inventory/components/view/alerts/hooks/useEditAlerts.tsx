import { ToastProps } from '@components/toast/Toast';
import { FormEvent, useState } from 'react';
import settingsService from '../../../../../../services/settingsService';
import { AlertMethod, Alert } from './useAlerts';

type AlertType = 'BUDGET' | 'USAGE';

const ALERT_TYPE = {
  BUDGET: 'BUDGET',
  USAGE: 'USAGE'
} as const;

type Options = {
  label: 'Cost' | 'Resources';
  image: string;
  description: string;
  type: AlertType;
};

type useEditAlertsProps = {
  alertMethod: AlertMethod;
  currentAlert: Alert | undefined;
  viewId: number;
  closeAlert: (action?: 'hasChanges' | undefined) => void;
  showToast: (toast: ToastProps) => void;
};

const INITIAL_BUDGET_ALERT: Partial<Alert> = {
  viewId: '',
  name: '',
  type: 'BUDGET',
  budget: '0'
};

const INITIAL_USAGE_ALERT: Partial<Alert> = {
  viewId: '',
  name: '',
  type: 'USAGE',
  usage: '0'
};

function useEditAlerts({
  alertMethod: alertType,
  viewId,
  currentAlert,
  closeAlert,
  showToast
}: useEditAlertsProps) {
  const [selected, setSelected] = useState<AlertType>(
    currentAlert?.type || ALERT_TYPE.BUDGET
  );
  const [alert, setAlert] = useState<Partial<Alert>>(
    currentAlert || INITIAL_BUDGET_ALERT
  );
  const [loading, setLoading] = useState(false);

  const options: Options[] = [
    {
      label: 'Cost',
      image: '/assets/img/others/cost.svg',
      description: 'If the total cost goes over the limit threshold',
      type: 'BUDGET'
    },
    {
      label: 'Resources',
      image: '/assets/img/others/resource.svg',
      description: 'If the number of resources goes over the limit',
      type: 'USAGE'
    }
  ];

  function changeAlertType(type: AlertType) {
    if (type === ALERT_TYPE.BUDGET) {
      setAlert(INITIAL_BUDGET_ALERT);
      setSelected(type);
    }

    if (type === ALERT_TYPE.USAGE) {
      setAlert(INITIAL_USAGE_ALERT);
      setSelected(type);
    }
  }

  function handleChange(newData: Partial<Alert>) {
    setAlert(prev => ({ ...prev, ...newData }));
  }

  function submit(
    e: FormEvent<HTMLFormElement>,
    setViewControllerToAlertsBase: () => void,
    edit?: 'edit'
  ) {
    e.preventDefault();
    setLoading(true);

    const payload = { ...alert };

    if (payload.type === ALERT_TYPE.BUDGET) {
      payload.budget = Number(payload.budget);
    }

    if (payload.type === ALERT_TYPE.USAGE) {
      payload.usage = Number(payload.usage);
    }

    payload.isSlack = alertType === 0;
    if (!edit) {
      payload.viewId = viewId.toString();
      const payloadJson = JSON.stringify(payload);
      settingsService.createAlert(payloadJson).then(res => {
        if (res === Error || res.error) {
          setLoading(false);
          showToast({
            hasError: true,
            title: 'Alert not created',
            message:
              'There was an error creating this alert. Refer to the logs and try again.'
          });
        } else {
          setLoading(false);
          showToast({
            hasError: false,
            title: 'Alert created',
            message: `The alert was successfully created!`
          });
          closeAlert('hasChanges');
          setViewControllerToAlertsBase();
        }
      });
    }

    if (edit) {
      const { id } = payload;

      if (id) {
        const payloadJson = JSON.stringify(payload);
        settingsService.editAlert(id, payloadJson).then(res => {
          if (res === Error || res.error) {
            setLoading(false);
            showToast({
              hasError: true,
              title: 'Alert not edited',
              message:
                'There was an error editing this alert. Refer to the logs and try again.'
            });
          } else {
            setLoading(false);
            showToast({
              hasError: false,
              title: 'Alert edited',
              message: `The alert was successfully edited!`
            });
            closeAlert('hasChanges');
          }
        });
      }
    }
  }

  function deleteAlert(alertId: number) {
    const id = alertId;

    settingsService.deleteAlert(id).then(res => {
      if (res === Error || res.error) {
        setLoading(false);
        showToast({
          hasError: true,
          title: 'Alert was not deleted',
          message:
            'There was an error deleting this alert. Refer to the logs and try again.'
        });
      } else {
        setLoading(false);
        showToast({
          hasError: false,
          title: 'Alert deleted',
          message: `The alert was successfully deleted!`
        });
        closeAlert('hasChanges');
      }
    });
  }

  const buttonDisabled = !alert.name || (!alert.budget && !alert.usage);

  return {
    selected,
    options,
    alert,
    changeAlertType,
    handleChange,
    buttonDisabled,
    submit,
    loading,
    deleteAlert
  };
}

export default useEditAlerts;
