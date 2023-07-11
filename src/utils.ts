export function formatDate(dateString: string): string {

    const date = new Date(dateString);

    const opciones: Intl.DateTimeFormatOptions = { 
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: 'numeric',
      minute: 'numeric',
      hour12: false,
    };
     
    return date.toLocaleDateString('en-US', opciones);

}

export function shortFormatDate(dateString: string): string {
      
    const date = new Date(dateString);

    const opciones: Intl.DateTimeFormatOptions = {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    };
    
    return date.toLocaleDateString('en-US', opciones);
    
}