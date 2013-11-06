var jsonrpc = {
    version: 0.1,
    maxRequest: 10
};

jsonrpc.Call = function(request){
    while( this.request[this.i] != null ) {
        this.i++;
        if(this.i >= this.maxRequest) {
            this.i = 0;
        };
    }

    request.id = this.i;
    this.request[this.i] = request;
    data = JSON.stringify(request);
    this.ws.send(data);
};

jsonrpc.onclientmessage = function(e) {
    var client = e.originalTarget.parrent;
    var d = e.data;
    ret = JSON.parse(e.data, function (key, value) {
        var type;
        if (value && typeof value === 'object') {
            type = value.type;
            if (typeof type === 'string' && typeof window[type] === 'function') {
                return new (window[type])(value);
            }
        }
        return value;
    });

    if(ret.error == null) {
        client.request[ret.id].success(ret.result);
    } else {
        client.request[ret.id].error(ret.error);
    };
    client.request[ret.id] = null;
};

jsonrpc.Close = function(){
    this.ws.Close;
};

jsonrpc.NewClient = function(addr){
    var client = {};
    client.i = 0;
    client.maxRequest = jsonrpc.maxRequest;
    client.request = new Array();
    client.Addr = addr;
    client.ws = new WebSocket(addr);
    client.ws.parrent = client;
    client.ws.onmessage = jsonrpc.onclientmessage;
    client.Call = jsonrpc.Call;
    client.Close = jsonrpc.Close;
    return client;
};
