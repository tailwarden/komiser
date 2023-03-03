import { FormEvent, useState } from 'react';
import settingsService from '../../../../../../services/settingsService';
import useToast from '../../../../../toast/hooks/useToast';
import { SlackAlert } from './useSlackAlerts';

type SlackAlertType = 'BUDGET' | 'USAGE';

type Options = {
  label: 'Cost' | 'Resources';
  description: string;
  type: SlackAlertType;
};

type useEditSlackAlertsProps = {
  currentSlackAlert: SlackAlert | undefined;
  viewId: number;
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
  currentSlackAlert
}: useEditSlackAlertsProps) {
  const [selected, setSelected] = useState<SlackAlertType>(
    currentSlackAlert?.type || 'BUDGET'
  );
  const [slackAlert, setSlackAlert] = useState<Partial<SlackAlert>>(
    currentSlackAlert || INITIAL_BUDGET_SLACK_ALERT
  );
  const [loading, setLoading] = useState(false);
  const { toast, setToast, dismissToast } = useToast();

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
    if (type === 'BUDGET') {
      setSlackAlert(INITIAL_BUDGET_SLACK_ALERT);
      setSelected(type);
    }

    if (type === 'USAGE') {
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

    if (payload.type === 'BUDGET') {
      payload.budget = Number(payload.budget);
    }

    if (payload.type === 'USAGE') {
      payload.usage = Number(payload.usage);
    }

    if (!edit) {
      payload.viewId = viewId.toString();
      const payloadJson = JSON.stringify(payload);
      settingsService.createSlackAlert(payloadJson).then(res => {
        if (res === Error) {
          setLoading(false);
          setToast({
            hasError: false,
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
        }
      });
    }

    if (edit) {
      const id = payload.id?.toString();

      if (id) {
        const payloadJson = JSON.stringify(payload);
        settingsService.editSlackAlert(id, payloadJson).then(res => {
          if (res === Error) {
            setLoading(false);
            setToast({
              hasError: false,
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
          }
        });
      }
    }
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
    loading
  };
}

export default useEditSlackAlerts;
