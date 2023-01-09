export namespace examples {
	
	export class Example {
	    title: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new Example(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.content = source["content"];
	    }
	}

}

