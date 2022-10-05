import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
    name: 'search',
})
export class SearchFilterPipe implements PipeTransform {
    transform(services: any, term: any): any {
        if (!term) return services;

        return services.filter((service) => {
            return (
                service.region.toLowerCase().includes(term.toLowerCase()) ||
                service.account.toLowerCase().includes(term.toLowerCase()) ||
                service.provider.toLowerCase().includes(term.toLowerCase()) ||
                service.service.toLowerCase().includes(term.toLowerCase()) ||
                service.name.toLowerCase().includes(term.toLowerCase()) || 
                (service.tags?.filter((tag) => {
                    tag.toLowerCase().includes(term.toLowerCase())
                }).length > 0)
            );
        });
    }
}