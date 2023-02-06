import environment from '../environments/environment';

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
  async getInventoryStats() {
    try {
      const res = await fetch(`${BASE_URL}/stats`, settings('GET'));
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getInventoryList(urlParams: string) {
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

  async searchInventory(query: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/inventory?query=${query}`,
        settings('GET')
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

  async getFilteredInventory(urlParams: string, payload: string) {
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

  async getCustomViewInventory(id: string, urlParams: string, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/views/${id}/resources${urlParams}`,
        settings('POST', payload)
      );
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
  }
};

export default settingsService;
