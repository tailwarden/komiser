import { useEffect, useState } from 'react';
import settingsService from '../../../../../../services/settingsService';

type useSlackAlertsProps = {
  viewId: number;
};

export type SlackAlert = {
  id: number;
  name: string;
  viewId: string;
  type: 'BUDGET' | 'USAGE';
  budget?: number | string;
  usage?: number | string;
};

function useSlackAlerts({ viewId }: useSlackAlertsProps) {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(true);
  const [hasSlack, setHasSlack] = useState(false);
  const [slackAlerts, setSlackAlerts] = useState<SlackAlert[]>();
  const [editSlackAlert, setEditSlackAlert] = useState(false);
  const [currentSlackAlert, setCurrentSlackAlert] = useState<SlackAlert>();
  const currentViewId = viewId.toString();

  function fetchSlackStatus() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getSlackIntegration().then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setHasSlack(res.enabled);
      }
    });
  }

  function fetchViewAlerts() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getSlackAlertsFromAView(currentViewId).then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setSlackAlerts(res);
      }
    });
  }

  function createOrEditSlackAlert(alertId?: number) {
    if (alertId && slackAlerts) {
      const foundSlackAlert = slackAlerts.find(alert => alert.id === alertId);

      if (foundSlackAlert) {
        setCurrentSlackAlert(foundSlackAlert);
      }
    }
    setEditSlackAlert(true);
  }

  function closeSlackAlert(action?: 'hasChanges') {
    setCurrentSlackAlert(undefined);
    setEditSlackAlert(false);

    if (action === 'hasChanges') {
      fetchViewAlerts();
    }
  }

  useEffect(() => {
    if (!hasSlack) {
      fetchSlackStatus();
    }

    if (hasSlack && currentViewId) {
      fetchViewAlerts();
    }
  }, [hasSlack]);

  const hasNoSlackAlerts =
    hasSlack && !editSlackAlert && slackAlerts && slackAlerts.length === 0;

  return {
    loading,
    error,
    hasSlack,
    slackAlerts,
    hasNoSlackAlerts,
    editSlackAlert,
    currentSlackAlert,
    createOrEditSlackAlert,
    closeSlackAlert,
    fetchViewAlerts
  };
}

export default useSlackAlerts;
