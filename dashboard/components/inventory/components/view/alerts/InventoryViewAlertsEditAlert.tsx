import classNames from 'classnames';
import { useState } from 'react';
import Image from 'next/image';
import { ToastProps } from '@components/toast/Toast';
import Button from '../../../../button/Button';
import Grid from '../../../../grid/Grid';
import ArrowLeftIcon from '../../../../icons/ArrowLeftIcon';
import Input from '../../../../input/Input';
import useEditAlerts from './hooks/useEditAlerts';
import { AlertMethod, Alert } from './hooks/useAlerts';
import settingsService from '../../../../../services/settingsService';
import LoadingSpinner from '../../../../icons/LoadingSpinner';

type InventoryViewAlertsCreateOrEditAlertProps = {
  alertMethod: AlertMethod;
  setViewControllerOnSubmit: () => void;
  setViewControllerOnClickingBackButton: () => void;
  setViewControllerOnDelete: () => void;
  currentAlert: Alert | undefined;
  closeAlert: (action?: 'hasChanges' | undefined) => void;
  viewId: number;
  showToast: (toast: ToastProps) => void;
};

function InventoryViewAlertsCreateOrEditAlert({
  alertMethod,
  setViewControllerOnSubmit,
  setViewControllerOnClickingBackButton,
  setViewControllerOnDelete,
  currentAlert,
  closeAlert,
  viewId,
  showToast
}: InventoryViewAlertsCreateOrEditAlertProps) {
  const {
    selected,
    options,
    alert,
    changeAlertType,
    handleChange,
    buttonDisabled,
    submit,
    loading
  } = useEditAlerts({
    alertMethod,
    currentAlert,
    viewId,
    closeAlert,
    showToast
  });

  const [testingEndpoint, setTestingEndpoint] = useState(false);
  const [testResultData, setTestResultData] = useState<{
    success: boolean;
    message: string;
  }>();
  const [testResultSuccess, setTestResultSuccess] = useState<boolean>(false);

  const findWhichOption =
    currentAlert && options.find(option => option.type === currentAlert.type);

  let alertName = alertMethod === 0 ? 'Slack' : 'Webhook';
  if (!currentAlert) {
    alert.isSlack = alertName !== 'Webhook';
  } else {
    alert.isSlack = currentAlert.isSlack;
    alertName = currentAlert.isSlack ? 'Slack' : 'Webhook';
  }

  const testEndpoint = async () => {
    if (alert.endpoint) {
      setTestingEndpoint(true);
      settingsService.testEndpoint(alert.endpoint).then(data => {
        setTestingEndpoint(false);
        setTestResultSuccess(data.success);
        setTestResultData({ success: data.success, message: data.message });

        setTimeout(() => {
          setTestResultSuccess(false);
        }, 3000);
      });
    } else {
      setTestResultData({
        success: false,
        message: 'Please type an endpoint above'
      });
    }
  };

  return (
    <form
      onSubmit={e => {
        if (currentAlert) {
          submit(e, setViewControllerOnSubmit, 'edit');
        } else {
          submit(e, setViewControllerOnSubmit);
        }
      }}
      className="flex flex-col gap-6 text-sm"
    >
      {!currentAlert && (
        <div
          onClick={() => setViewControllerOnClickingBackButton()}
          className="flex cursor-pointer items-center gap-2 self-start text-gray-950"
        >
          <ArrowLeftIcon width={24} height={24} />
          Setup {alertName} Alert
        </div>
      )}
      {/* Display a back button if editing a Slack alert */}
      {currentAlert && (
        <div
          onClick={() => closeAlert()}
          className="flex cursor-pointer items-center gap-2 self-start text-gray-950"
        >
          <ArrowLeftIcon width={24} height={24} />
          Edit {currentAlert.isSlack ? 'Slack' : 'Webhook'} alert
        </div>
      )}

      <div className="flex flex-col gap-4">
        <p className="text-gray-700">Type</p>

        {/* Displaying the slack alert types when creating a new alert */}
        {!currentAlert && (
          <Grid gap="sm">
            {options.map(option => {
              const isActive = selected === option.type;
              return (
                <div
                  key={option.label}
                  onClick={() => changeAlertType(option.type)}
                  className={classNames(
                    'flex cursor-pointer flex-col items-start justify-center rounded-lg px-6 py-4 outline outline-gray-300 hover:outline-gray-500',
                    {
                      'outline-2 outline-darkcyan-500 hover:outline-darkcyan-500':
                        isActive
                    }
                  )}
                >
                  <div className="flex items-center gap-4">
                    <Image
                      src={option.image}
                      alt={option.label}
                      height={42}
                      width={42}
                    />
                    <div className="flex flex-col">
                      <p className="text-base font-semibold text-gray-950">
                        {option.label}
                      </p>
                      <p className="text-xs text-gray-700">
                        {option.description}
                      </p>
                    </div>
                  </div>
                </div>
              );
            })}
          </Grid>
        )}

        {/* Displaying the chosen slack alert type when editing an alert */}
        {currentAlert && (
          <div className="flex flex-col items-start justify-center">
            <p className="text-base font-semibold text-gray-950">
              {findWhichOption?.label}
            </p>
            <p className="text-xs text-gray-700">
              {findWhichOption?.description}
            </p>
          </div>
        )}
      </div>
      <div className="flex flex-col gap-4">
        <p className="text-gray-700">Details</p>

        {selected === 'BUDGET' && (
          <Grid gap="sm">
            <Input
              type="text"
              label="Name"
              name="name"
              action={handleChange}
              value={alert.name}
              maxLength={64}
            />
            <Input
              type="number"
              label="Limit (in $)"
              name="budget"
              action={handleChange}
              value={alert.budget}
              min={0}
              positiveNumberOnly
            />
          </Grid>
        )}

        {selected === 'USAGE' && (
          <Grid gap="sm">
            <Input
              type="text"
              label="Name"
              name="name"
              action={handleChange}
              value={alert.name}
            />
            <Input
              type="number"
              label="Limit (of resources)"
              name="usage"
              action={handleChange}
              value={alert.usage}
              min={0}
              positiveNumberOnly
            />
          </Grid>
        )}
      </div>
      {alertName === 'Webhook' && (
        <div>
          <div className="flex flex-col gap-4">
            <p className="text-gray-700">Output</p>
            <div className="relative">
              <div>
                <div className="relative">
                  <input
                    type="text"
                    name="endpoint"
                    className={`peer mr-6 w-full rounded bg-white pb-[0.75rem] pl-4 pr-32 pt-[1.75rem] text-sm text-gray-950 caret-darkcyan-500 outline outline-gray-300 focus:outline-2 focus:outline-darkcyan-500 ${
                      testResultData?.success === false &&
                      `outline-red-500 focus:outline-red-500`
                    }`}
                    placeholder=""
                    onChange={e => {
                      handleChange({ endpoint: e.target.value });
                    }}
                    value={alert.endpoint}
                    maxLength={64}
                    autoComplete="off"
                    data-lpignore="true"
                    data-form-type="other"
                  />
                  <span className="pointer-events-none absolute bottom-[1.925rem] left-4 origin-left scale-75 select-none font-normal text-gray-500 transition-all peer-placeholder-shown:bottom-[1.15rem] peer-placeholder-shown:left-4 peer-placeholder-shown:scale-[87.5%] peer-focus:bottom-[1.925rem] peer-focus:scale-75">
                    Endpoint
                  </span>
                </div>
              </div>
              <span
                className={`absolute right-4 top-1/2 flex w-full -translate-y-1/2 transform cursor-pointer items-center gap-2 rounded bg-transparent text-sm font-medium text-darkcyan-500 text-darkcyan-500 active:bg-cyan-200 active:text-darkcyan-500 disabled:cursor-not-allowed sm:w-auto ${
                  testingEndpoint ? 'pointer-events-none opacity-50' : ''
                }`}
                onClick={testEndpoint}
              >
                {testingEndpoint ? (
                  <>
                    <LoadingSpinner />
                    <span className="align-center ml-1 flex justify-center">
                      Test Endpoint
                    </span>
                  </>
                ) : (
                  <>
                    {testResultSuccess && (
                      <Image
                        src="/assets/img/others/tickmark.svg"
                        height={20}
                        width={20}
                        alt="tickmark"
                        className="-mr-1"
                      />
                    )}

                    <span className={testResultSuccess ? 'text-green-600' : ''}>
                      {testResultSuccess ? 'Valid Endpoint' : 'Test Endpoint'}
                    </span>
                  </>
                )}
              </span>
            </div>
            {testResultData?.success === false && (
              <p className="-mt-2 text-xs text-red-500">
                {testResultData.message}
              </p>
            )}

            <Input
              type="text"
              label="Secret (optional)"
              name="secret"
              action={handleChange}
              value={alert.secret}
            />
          </div>
          <div className="mt-2">
            <p className="text-xs text-gray-700">
              We’ll send a POST request to the endpoint. More information can be
              found in our{' '}
              <a
                href="https://docs.komiser.io/docs/guides/how-to-komiser/alerts#request-details"
                target="_blank"
                rel="noreferrer"
                className="text-darkcyan-500"
              >
                <u>developer documentation</u>
              </a>
              .
            </p>
          </div>
        </div>
      )}

      <div className="flex items-center justify-between">
        <div>
          {/* Display a delete button if it's editing an alert */}
          {currentAlert && (
            <Button
              size="lg"
              style="delete"
              onClick={setViewControllerOnDelete}
            >
              Delete alert
            </Button>
          )}
        </div>

        <div className="flex gap-4">
          <Button style="ghost" size="lg" onClick={closeAlert}>
            Cancel
          </Button>
          <Button
            type="submit"
            size="lg"
            disabled={buttonDisabled}
            loading={loading}
          >
            {currentAlert ? 'Save changes' : 'Set up alert'}
          </Button>
        </div>
      </div>
    </form>
  );
}

export default InventoryViewAlertsCreateOrEditAlert;
