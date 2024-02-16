import environment from '../environments/environment';
import { InventoryFilterData } from '../components/inventory/hooks/useInventory/types/useInventoryTypes';

const BASE_URL = environment.API_URL;

type Settings = {
  method: 'GET' | 'PUT' | 'POST' | 'DELETE' | undefined;
  headers: {
    'Content-Type': 'application/json';
  };
  body?: string;
};

function settings(arg: string, payload?: string) {
  const settingsObj: Settings = {
    method: undefined,
    headers: {
      'Content-Type': 'application/json'
    },
    body: undefined
  };

  if (arg === 'GET') {
    settingsObj.method = 'GET';
    delete settingsObj.body;
  }

  if (arg === 'PUT') {
    settingsObj.method = 'PUT';
    delete settingsObj.body;
  }

  if (arg === 'PUT' && payload) {
    settingsObj.method = 'PUT';
    settingsObj.body = payload;
  }

  if (arg === 'POST') {
    settingsObj.method = 'POST';
    delete settingsObj.body;
  }

  if (arg === 'POST' && payload) {
    settingsObj.method = 'POST';
    settingsObj.body = payload;
  }

  if (arg === 'DELETE') {
    settingsObj.method = 'DELETE';
  }

  return settingsObj;
}

const settingsService = {
  async getGlobalStats() {
    try {
      const res = await fetch(`${BASE_URL}/global/stats`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getTelemetry() {
    try {
      const res = await fetch(`${BASE_URL}/telemetry`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getGlobalLocations() {
    try {
      const res = await fetch(`${BASE_URL}/global/locations`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getRelations(payload: InventoryFilterData[]) {
    try {
      const res = await fetch(
        `${BASE_URL}/resources/relations`,
        settings('POST', JSON.stringify(payload))
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getGlobalResources(payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/global/resources`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getCostExplorer(payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/costs/explorer`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getInventory(urlParams: string, payload: string = '[]') {
    try {
      const res = await fetch(
        `${BASE_URL}/resources/search${urlParams}`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getResourceById(urlParams: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/resources${urlParams}`,
        settings('GET')
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getInventoryStats() {
    try {
      const res = await fetch(`${BASE_URL}/stats`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getFilteredInventoryStats(payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/stats/search`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async saveTags(serviceId: string, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/resources/${serviceId}/tags`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async bulkSaveTags(payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/resources/tags`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getProviders() {
    try {
      const res = await fetch(`${BASE_URL}/providers`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getAccounts() {
    try {
      const res = await fetch(`${BASE_URL}/accounts`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getRegions() {
    try {
      const res = await fetch(`${BASE_URL}/regions`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getServices() {
    try {
      const res = await fetch(`${BASE_URL}/services`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getViews() {
    try {
      const res = await fetch(`${BASE_URL}/views`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getHiddenResourcesFromView(id: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/views/${id}/hidden/resources`,
        settings('GET')
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async hideResourceFromView(id: string, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/views/${id}/resources/hide`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async unhideResourceFromView(id: string, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/views/${id}/resources/unhide`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async saveView(payload: string) {
    try {
      const res = await fetch(`${BASE_URL}/views`, settings('POST', payload));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async updateView(id: string, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/views/${id}`,
        settings('PUT', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async deleteView(id: string) {
    try {
      const res = await fetch(`${BASE_URL}/views/${id}`, settings('DELETE'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getSlackIntegration() {
    try {
      const res = await fetch(`${BASE_URL}/slack`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async testEndpoint(endpoint: string) {
    if (!endpoint) return { success: false, message: 'Endpoint is required.' };
    try {
      const response = await fetch(`${BASE_URL}/alerts/test`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url: endpoint })
      });
      const data = await response.json();
      return data;
    } catch {
      return {
        success: false,
        message: 'Something went wrong!'
      };
    }
  },

  async getAlertsFromAView(id: number) {
    try {
      const res = await fetch(
        `${BASE_URL}/views/${id}/alerts`,
        settings('GET')
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async createAlert(payload: string) {
    try {
      const res = await fetch(`${BASE_URL}/alerts`, settings('POST', payload));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async editAlert(id: number, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/alerts/${id}`,
        settings('PUT', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async deleteAlert(id: number) {
    try {
      const res = await fetch(`${BASE_URL}/alerts/${id}`, settings('DELETE'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  exportCSV(id?: string) {
    return window.location.replace(
      `${BASE_URL}/resources/export-csv${id ? `/${id}` : ''}`
    );
  },

  async getCloudAccounts() {
    try {
      const res = await fetch(`${BASE_URL}/cloud_accounts`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async addCloudAccount(payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/cloud_accounts`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async editCloudAccount(id: number, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/cloud_accounts/${id}`,
        settings('PUT', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async deleteCloudAccount(id: number) {
    try {
      const res = await fetch(
        `${BASE_URL}/cloud_accounts/${id}`,
        settings('DELETE')
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async rescanCloudAccount(id: number) {
    try {
      const res = await fetch(
        `${BASE_URL}/cloud_accounts/resync/${id}`,
        settings('GET')
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async saveDatabaseConfig(payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/databases`,
        settings('POST', payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getOnboardingStatus() {
    try {
      const res = await fetch(`${BASE_URL}/is_onboarded`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async sendFeedback(payload: FormData) {
    try {
      const res = await fetch(`${BASE_URL}/feedback`, {
        method: 'POST',
        body: payload
      });
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  }
};

export default settingsService;
