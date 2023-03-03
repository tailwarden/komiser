import classNames from 'classnames';
import Button from '../../../../button/Button';
import Grid from '../../../../grid/Grid';
import Input from '../../../../input/Input';
import { ToastProps } from '../../../../toast/hooks/useToast';
import useEditSlackAlerts from './hooks/useEditSlackAlerts';
import { SlackAlert } from './hooks/useSlackAlerts';

type InventoryViewAlertsEditSlackAlertProps = {
  currentSlackAlert: SlackAlert | undefined;
  closeSlackAlert: (action?: 'hasChanges' | undefined) => void;
  viewId: number;
  setToast: (toast: ToastProps | undefined) => void;
};

function InventoryViewAlertsEditSlackAlert({
  currentSlackAlert,
  closeSlackAlert,
  viewId,
  setToast
}: InventoryViewAlertsEditSlackAlertProps) {
  const {
    selected,
    options,
    slackAlert,
    changeSlackAlertType,
    handleChange,
    buttonDisabled,
    submit,
    loading,
    deleteSlackAlert
  } = useEditSlackAlerts({
    currentSlackAlert,
    viewId,
    closeSlackAlert,
    setToast
  });

  const findWhichOption =
    currentSlackAlert &&
    options.find(option => option.type === currentSlackAlert.type);

  return (
    <form
      onSubmit={e => {
        if (currentSlackAlert) {
          submit(e, 'edit');
        } else {
          submit(e);
        }
      }}
      className="flex flex-col gap-6 text-sm"
    >
      {/* Display a back button if editing a Slack alert */}
      {currentSlackAlert && (
        <div
          onClick={() => closeSlackAlert()}
          className="flex cursor-pointer items-center gap-2 self-start text-black-900"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeMiterlimit="10"
              strokeWidth="1.5"
              d="M9.57 18.07L3.5 12l6.07-6.07M20.5 12H3.67"
            ></path>
          </svg>
          Edit alert
        </div>
      )}

      <div className="flex flex-col gap-4">
        <p className="text-black-400">Type</p>

        {/* Displaying the slack alert types when creating a new alert */}
        {!currentSlackAlert && (
          <Grid gap="sm">
            {options.map(option => {
              const isActive = selected === option.type;
              return (
                <div
                  key={option.label}
                  onClick={() => changeSlackAlertType(option.type)}
                  className={classNames(
                    'flex cursor-pointer flex-col items-start justify-center rounded-lg py-4 px-6 outline outline-black-200 hover:outline-black-300',
                    {
                      'outline-2 outline-primary hover:outline-primary':
                        isActive
                    }
                  )}
                >
                  <p className="text-base font-semibold text-black-900">
                    {option.label}
                  </p>
                  <p className="text-xs text-black-400">{option.description}</p>
                </div>
              );
            })}
          </Grid>
        )}

        {/* Displaying the chosen slack alert type when editing an alert */}
        {currentSlackAlert && (
          <div
            className={classNames('flex flex-col items-start justify-center')}
          >
            <p className="text-base font-semibold text-black-900">
              {findWhichOption?.label}
            </p>
            <p className="text-xs text-black-400">
              {findWhichOption?.description}
            </p>
          </div>
        )}
      </div>
      <div className="flex flex-col gap-4">
        <p className="text-black-400">Details</p>

        {selected === 'BUDGET' && (
          <Grid gap="sm">
            <Input
              label="Name"
              name="name"
              action={handleChange}
              value={slackAlert.name}
            />
            <Input
              type="number"
              label="Limit (in $)"
              name="budget"
              action={handleChange}
              value={slackAlert.budget}
              min={0}
              positiveNumberOnly
            />
          </Grid>
        )}

        {selected === 'USAGE' && (
          <Grid gap="sm">
            <Input
              label="Name"
              name="name"
              action={handleChange}
              value={slackAlert.name}
            />
            <Input
              type="number"
              label="Limit (of resources)"
              name="usage"
              action={handleChange}
              value={slackAlert.usage}
              min={0}
              positiveNumberOnly
            />
          </Grid>
        )}
      </div>

      <div className="flex items-center justify-between">
        <div>
          {/* Display a delete button if it's editing an alert */}
          {currentSlackAlert && (
            <Button
              size="lg"
              style="delete"
              onClick={() => deleteSlackAlert(currentSlackAlert.id)}
            >
              Delete alert
            </Button>
          )}
        </div>

        <div className="flex gap-4">
          <Button style="secondary" size="lg" onClick={closeSlackAlert}>
            Cancel
          </Button>
          <Button
            type="submit"
            size="lg"
            disabled={buttonDisabled}
            loading={loading}
          >
            {currentSlackAlert ? 'Save changes' : 'Set up alert'}
          </Button>
        </div>
      </div>
    </form>
  );
}

export default InventoryViewAlertsEditSlackAlert;
