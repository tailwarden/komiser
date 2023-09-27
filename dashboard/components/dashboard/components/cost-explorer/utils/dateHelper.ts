const dateHelper = {
  getToday() {
    const date = new Date();
    return date.toJSON().slice(0, 10);
  },

  getFirstDayOfThisMonth() {
    const date = new Date();
    date.setDate(1);
    return date.toJSON().slice(0, 10);
  },

  getLastMonth() {
    const date = new Date();
    date.setDate(1);
    date.setMonth(date.getMonth() - 1);
    return date.toJSON().slice(0, 10);
  },

  getLastDayOfLastMonth() {
    const date = new Date();
    date.setDate(0);
    date.setMonth(date.getMonth() + 1, 0);
    return date.toJSON().slice(0, 10);
  },

  getLastThreeMonths() {
    const date = new Date();
    date.setDate(1);
    date.setMonth(date.getMonth() - 3);
    return date.toJSON().slice(0, 10);
  },

  getLastSixMonths() {
    const date = new Date();
    date.setDate(1);
    date.setMonth(date.getMonth() - 6);
    return date.toJSON().slice(0, 10);
  },

  getLastTwelveMonths() {
    const date = new Date();
    date.setMonth(date.getMonth() - 12);
    return date.toJSON().slice(0, 10);
  },

  getMinDateStart() {
    const date = new Date();
    date.setFullYear(date.getFullYear() - 2);
    return date.toJSON().slice(0, 10);
  },

  getMaxDateStart() {
    return new Date();
  }
};

export function dateFormatter(dateParam: string, granularity: string) {
  const date = new Date(dateParam);
  let formattedDate;

  if (granularity === 'monthly') {
    formattedDate = date.toLocaleDateString('en-US', {
      timeZone: 'UTC',
      year: 'numeric',
      month: 'short'
    });
  }

  if (granularity === 'daily') {
    formattedDate = date.toLocaleDateString('en-US', {
      timeZone: 'UTC',
      year: 'numeric',
      month: 'short',
      day: '2-digit'
    });
  }
  return formattedDate;
}
export const thisMonth = [
  dateHelper.getFirstDayOfThisMonth(),
  dateHelper.getToday()
];
export const lastMonth = [
  dateHelper.getLastMonth(),
  dateHelper.getLastDayOfLastMonth()
];
export const lastThreeMonths = [
  dateHelper.getLastThreeMonths(),
  dateHelper.getLastDayOfLastMonth()
];
export const lastSixMonths = [
  dateHelper.getLastSixMonths(),
  dateHelper.getLastDayOfLastMonth()
];
export const lastTwelveMonths = [
  dateHelper.getLastTwelveMonths(),
  dateHelper.getLastDayOfLastMonth()
];

export default dateHelper;
