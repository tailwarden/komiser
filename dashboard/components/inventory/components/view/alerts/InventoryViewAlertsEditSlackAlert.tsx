import classNames from 'classnames';
import Button from '../../../../button/Button';
import Grid from '../../../../grid/Grid';
import Input from '../../../../input/Input';
import useEditSlackAlerts from './hooks/useEditSlackAlerts';

type InventoryViewAlertsEditSlackAlertProps = {
  closeSlackAlert: () => void;
  viewId: number;
};

function InventoryViewAlertsEditSlackAlert({
  closeSlackAlert,
  viewId
}: InventoryViewAlertsEditSlackAlertProps) {
  const {
    selected,
    options,
    slackAlert,
    changeSlackAlertType,
    handleChange,
    buttonDisabled,
    submit,
    loading
  } = useEditSlackAlerts({ viewId });

  return (
    <form onSubmit={e => submit(e)} className="flex flex-col gap-6 text-sm">
      <div className="flex flex-col gap-4">
        <p className="text-black-400">Type</p>
        <Grid gap="sm">
          {options.map(option => {
            const isActive = selected === option.label;
            return (
              <div
                key={option.label}
                onClick={() => changeSlackAlertType(option.type)}
                className={classNames(
                  'flex cursor-pointer flex-col items-start justify-center rounded-lg py-4 px-6 outline outline-black-200 hover:outline-black-300',
                  {
                    'outline-2 outline-primary hover:outline-primary': isActive
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
      </div>
      <div className="flex flex-col gap-4">
        <p className="text-black-400">Details</p>

        {selected === 'Cost' && (
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
            />
          </Grid>
        )}

        {selected === 'Resources' && (
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
            />
          </Grid>
        )}
      </div>
      <div className="flex gap-4 self-end">
        <Button style="secondary" size="lg" onClick={closeSlackAlert}>
          Cancel
        </Button>
        <Button
          type="submit"
          size="lg"
          disabled={buttonDisabled}
          loading={loading}
        >
          Set up alert
        </Button>
      </div>
    </form>
  );
}

export default InventoryViewAlertsEditSlackAlert;
