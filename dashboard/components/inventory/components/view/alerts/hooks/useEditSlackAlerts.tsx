import type { FormEvent } from 'react';
import { useState } from 'react';
import settingsService from '@/services/settingsService';
import type { ToastProps } from '@/components/toast/hooks/useToast';
import type { SlackAlert } from './useSlackAlerts';

type SlackAlertType = 'BUDGET' | 'USAGE';

const SLACK_ALERT_TYPE = {
  BUDGET: 'BUDGET',
  USAGE: 'USAGE'
} as const;

type Options = {
  label: 'Cost' | 'Resources';
  description: string;
  type: SlackAlertType;
};

type useEditSlackAlertsProps = {
  currentSlackAlert: SlackAlert | undefined;
  viewId: number;
  closeSlackAlert: (action?: 'hasChanges' | undefined) => void;
  setToast: (toast: ToastProps | undefined) => void;
};

const INITIAL_BUDGET_SLACK_ALERT: Partial<SlackAlert> = {
  viewId: '',
  name: '',
  type: 'BUDGET',
  budget: '0'
};

const INITIAL_USAGE_SLACK_ALERT: Partial<SlackAlert> = {
  viewId: '',
  name: '',
  type: 'USAGE',
  usage: '0'
};

function useEditSlackAlerts({
  viewId,
  currentSlackAlert,
  closeSlackAlert,
  setToast
}: useEditSlackAlertsProps) {
  const [selected, setSelected] = useState<SlackAlertType>(
    currentSlackAlert?.type || SLACK_ALERT_TYPE.BUDGET
  );
  const [slackAlert, setSlackAlert] = useState<Partial<SlackAlert>>(
    currentSlackAlert || INITIAL_BUDGET_SLACK_ALERT
  );
  const [loading, setLoading] = useState(false);

  const options: Options[] = [
    {
      label: 'Cost',
      description: 'If the total cost goes over the limit threshold',
      type: 'BUDGET'
    },
    {
      label: 'Resources',
      description: 'If the number of resources goes over the limit',
      type: 'USAGE'
    }
  ];

  function changeSlackAlertType(type: SlackAlertType) {
    if (type === SLACK_ALERT_TYPE.BUDGET) {
      setSlackAlert(INITIAL_BUDGET_SLACK_ALERT);
      setSelected(type);
    }

    if (type === SLACK_ALERT_TYPE.USAGE) {
      setSlackAlert(INITIAL_USAGE_SLACK_ALERT);
      setSelected(type);
    }
  }

  function handleChange(newData: Partial<SlackAlert>) {
    setSlackAlert(prev => ({ ...prev, ...newData }));
  }

  function submit(e: FormEvent<HTMLFormElement>, edit?: 'edit') {
    e.preventDefault();
    setLoading(true);

    const payload = { ...slackAlert };

    if (payload.type === SLACK_ALERT_TYPE.BUDGET) {
      payload.budget = Number(payload.budget);
    }

    if (payload.type === SLACK_ALERT_TYPE.USAGE) {
      payload.usage = Number(payload.usage);
    }

    if (!edit) {
      payload.viewId = viewId.toString();
      const payloadJson = JSON.stringify(payload);
      settingsService.createSlackAlert(payloadJson).then(res => {
        if (res === Error || res.error) {
          setLoading(false);
          setToast({
            hasError: true,
            title: 'Alert not created',
            message:
              'There was an error creating this slack alert. Refer to the logs and try again.'
          });
        } else {
          setLoading(false);
          setToast({
            hasError: false,
            title: 'Alert created',
            message: `The slack alert was successfully created!`
          });
          closeSlackAlert('hasChanges');
        }
      });
    }

    if (edit) {
      const { id } = payload;

      if (id) {
        const payloadJson = JSON.stringify(payload);
        settingsService.editSlackAlert(id, payloadJson).then(res => {
          if (res === Error || res.error) {
            setLoading(false);
            setToast({
              hasError: true,
              title: 'Alert not edited',
              message:
                'There was an error editing this slack alert. Refer to the logs and try again.'
            });
          } else {
            setLoading(false);
            setToast({
              hasError: false,
              title: 'Alert edited',
              message: `The slack alert was successfully edited!`
            });
            closeSlackAlert('hasChanges');
          }
        });
      }
    }
  }

  function deleteSlackAlert(alertId: number) {
    const id = alertId;

    settingsService.deleteSlackAlert(id).then(res => {
      if (res === Error || res.error) {
        setLoading(false);
        setToast({
          hasError: true,
          title: 'Alert was not deleted',
          message:
            'There was an error deleting this slack alert. Refer to the logs and try again.'
        });
      } else {
        setLoading(false);
        setToast({
          hasError: false,
          title: 'Alert deleted',
          message: `The slack alert was successfully deleted!`
        });
        closeSlackAlert('hasChanges');
      }
    });
  }

  const buttonDisabled =
    !slackAlert.name || (!slackAlert.budget && !slackAlert.usage);

  return {
    selected,
    options,
    slackAlert,
    changeSlackAlertType,
    handleChange,
    buttonDisabled,
    submit,
    loading,
    deleteSlackAlert
  };
}

export default useEditSlackAlerts;
