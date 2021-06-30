using Pulumi;
using Minio = Pulumi.Minio;

class MyStack : Stack
{
    public MyStack()
    {
        var user = new Minio.IamUser("csharp-user");

        this.UserName = user.Name;
    }

    [Output] public Output<string> UserName { get; set; }
}
