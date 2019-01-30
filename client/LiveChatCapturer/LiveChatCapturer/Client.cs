using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Runtime.Serialization;
using System.Net.Http;
using System.Windows.Forms;
using Newtonsoft.Json;

using System.IO;
using System.Collections;

using System.Text;
using System.Threading;
using System.Text.RegularExpressions;

namespace LiveChatCapturer
{
    enum Status
    {
        Start,Connect,RequestLivePage,RequestLiveChat,UpdateCommentBox
    }
    class LiveChatMessage
    {
        public string sender;
        public string message;
        public string purchase;
    }
    class LiveChatResponse
    {
        public string continuation;
        public LiveChatMessage[] messages;
    }
    class ConnectResponse
    {
        public string status;
        public string continuation;
    }
    class DisplayForm
    {
        public Form1 Form { get; }
        public TextBox CommentBox { get; }
        public DisplayForm()
        {
            Form = new Form1();
            Form.ControlBox = false;
            Form.Text = "";
            CommentBox = new TextBox();
            CommentBox.Multiline = true;
            CommentBox.Width = Form.Width;
            CommentBox.Height = Form.Height;
            Form.Resize += (_, __) =>
            {
                CommentBox.Width = Form.Width;
                CommentBox.Height = Form.Height;
            };
            Form.Controls.Add(CommentBox);
        }
    }
    class Client
    {
        private HttpClient client;
        private Uri uri;
        private TextBox commentBox;
        private string continuation;
        private ArrayList filter;
        private string livePage;
        public Client(string configFile,TextBox box)
        {
            client = new HttpClient();
            commentBox = box;
            string server;
            string port;
            var filterStrings = new ArrayList();
            using (var reader = new StreamReader(configFile))
            {
                livePage = reader.ReadLine();
                server = reader.ReadLine();
                port = reader.ReadLine();
                string s;
                while ((s = reader.ReadLine()) != null)
                {
                    filterStrings.Add(s);
                }
            }
            var builder = new UriBuilder(server);
            builder.Port = Convert.ToInt32(port);
            uri = builder.Uri;
            filter = new ArrayList();
            foreach (string filterString in filterStrings)
            {
                filter.Add(new Regex(filterString));
            }
        }
        public void ConnectToLivePage()
        {
            var request = new HttpRequestMessage(HttpMethod.Get,uri);
            request.Headers.Add("Action", "Connect");
            request.Headers.Add("Live-Page", livePage);
            var response = client.SendAsync(request).Result;
            var data = response.Content.ReadAsStringAsync().Result;
            var status = JsonConvert.DeserializeObject<ConnectResponse>(data);
            if (status.status != null && status.status == "Success" && status.continuation != null && status.continuation.Length>0) continuation = status.continuation;
            else {
                MessageBox.Show("Connect To Live Page Error.");
                System.Environment.Exit(-1);
            }
        }
        public void UpdateComment()
        {
            var request = new HttpRequestMessage(HttpMethod.Get, uri);
            request.Headers.Add("Action", "Update");
            request.Headers.Add("Continuation", continuation);
            var response = client.SendAsync(request).Result.Content.ReadAsStringAsync().Result;
            var resp = JsonConvert.DeserializeObject<LiveChatResponse>(response);
            if (resp.continuation != null && resp.continuation.Length > 0)
            {
                continuation = resp.continuation;
            }
            if (resp.messages != null)
            {
                foreach (var message in resp.messages)
                {
                    bool isBlock = false;
                    foreach (Regex reg in filter)
                    {
                        if (reg.Matches(message.message).Count != 0)
                        {
                            isBlock = true;
                        }
                    }
                    if (!isBlock)
                    {
                        string msg = "";
                        if (message.purchase != null && message.purchase != "")
                        {
                            msg += "[SuperChat(" + message.purchase + ")]";
                        }
                        msg += message.sender + ": " + message.message + "\r\n";
                        commentBox.BeginInvoke(new Action(() => commentBox.AppendText(msg)));
                    }
                }
            }
        }
        public void Start()
        {
            ConnectToLivePage();
            while (true)
            {
                UpdateComment();
                Thread.Sleep(500);
            } 
        }
    }
}
