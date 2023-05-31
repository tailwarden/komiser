import { useEffect, useState } from 'react';
import settingsService from '../../../../../../services/settingsService';

type useAlertsProps = {
  viewId: number;
};

export type Alert = {
  id: number;
  name: string;
  viewId: string;
  type: 'BUDGET' | 'USAGE';
  budget?: number | string;
  usage?: number | string;
  isSlack: boolean;
  endpoint?: string;
  secret?: string;
};

// eslint-disable-next-line no-shadow
export enum AlertMethod {
  'SLACK',
  'WEBHOOK'
}

function useAlerts({ viewId }: useAlertsProps) {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);
  const [isSlackConfigured, setIsSlackConfigured] = useState(false);
  const [hasAlerts, setHasAlerts] = useState(false);
  const [alerts, setAlerts] = useState<Alert[]>();
  const [editAlert, setEditAlert] = useState(false);
  const [currentAlert, setCurrentAlert] = useState<Alert>();
  const [alertsViewController, setAlertsViewController] = useState(0);
  const [alertMethod, setAlertMethod] = useState<AlertMethod>(
    AlertMethod.SLACK
  );

  function fetchAlertStatus() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getAlertsFromAView(viewId).then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setHasAlerts(res.length > 0);
      }
    });

    settingsService.getSlackIntegration().then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setIsSlackConfigured(res.enabled);
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

    settingsService.getAlertsFromAView(viewId).then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setHasAlerts(res.length > 0);
        setAlerts(res);
      }
    });
  }

  function incrementViewController() {
    setAlertsViewController(alertsViewController + 1);
  }

  function decrementViewController() {
    setAlertsViewController(alertsViewController - 1);
  }

  function setViewControllerToAlertsBaseView() {
    setAlertsViewController(0);
  }

  function setViewControllerToDeleteView() {
    setEditAlert(false);
    setAlertsViewController(3);
  }

  function setAlertMethodInAndIncrementViewController(alertName: AlertMethod) {
    incrementViewController();
    setAlertMethod(alertName);
  }

  function createOrEditAlert(alertId?: number) {
    if (alertId && alerts) {
      const foundAlert = alerts.find(alert => alert.id === alertId);

      if (foundAlert) {
        setCurrentAlert(foundAlert);
      }
    }
    setEditAlert(true);
  }

  function closeAlert(action?: 'hasChanges') {
    setCurrentAlert(undefined);
    setEditAlert(false);
    setViewControllerToAlertsBaseView();

    if (action === 'hasChanges') {
      fetchViewAlerts();
    }
  }

  useEffect(() => {
    if (!hasAlerts) {
      fetchAlertStatus();
    }

    if (hasAlerts && viewId) {
      fetchViewAlerts();
    }
  }, [hasAlerts]);

  return {
    loading,
    error,
    hasAlerts,
    isSlackConfigured,
    alerts,
    editAlert,
    currentAlert,
    alertsViewController,
    alertMethod,
    createOrEditAlert,
    setViewControllerToAlertsBaseView,
    setViewControllerToDeleteView,
    closeAlert,
    fetchViewAlerts,
    setAlertMethodInAndIncrementViewController,
    decrementViewController,
    incrementViewController
  };
}

export default useAlerts;
