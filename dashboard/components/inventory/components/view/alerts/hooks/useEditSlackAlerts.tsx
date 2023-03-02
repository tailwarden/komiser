import { FormEvent, useState } from 'react';
import settingsService from '../../../../../../services/settingsService';
import useToast from '../../../../../toast/hooks/useToast';

type Selected = 'Cost' | 'Resources';

type SlackAlertType = 'BUDGET' | 'USAGE';

type Options = {
  label: Selected;
  description: string;
  type: SlackAlertType;
};

type SlackAlert = {
  name: string;
  viewId: string;
  type: SlackAlertType;
  budget?: number | string;
  usage?: number | string;
};

type useEditSlackAlertsProps = {
  viewId: number;
};

const INITIAL_BUDGET_SLACK_ALERT: SlackAlert = {
  viewId: '',
  name: '',
  type: 'BUDGET',
  budget: '0'
};

const INITIAL_USAGE_SLACK_ALERT: SlackAlert = {
  viewId: '',
  name: '',
  type: 'USAGE',
  usage: '0'
};

function useEditSlackAlerts({ viewId }: useEditSlackAlertsProps) {
  const [selected, setSelected] = useState<Selected>('Cost');
  const [slackAlert, setSlackAlert] = useState<SlackAlert>(
    INITIAL_BUDGET_SLACK_ALERT
  );
  const [loading, setLoading] = useState(false);
  const { setToast } = useToast();

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
      setSelected('Cost');
    }

    if (type === 'USAGE') {
      setSlackAlert(INITIAL_USAGE_SLACK_ALERT);
      setSelected('Resources');
    }
  }

  function handleChange(newData: Partial<SlackAlert>) {
    setSlackAlert(prev => ({ ...prev, ...newData }));
  }

  function submit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setLoading(true);

    const payload = { ...slackAlert };
    payload.viewId = viewId.toString();

    if (payload.type === 'BUDGET') {
      payload.budget = Number(payload.budget);
    }

    if (payload.type === 'USAGE') {
      payload.usage = Number(payload.usage);
    }

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
