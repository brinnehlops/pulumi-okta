// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App
{
    /// <summary>
    /// Creates a Bookmark Application.
    /// 
    /// This resource allows you to create and configure a Bookmark Application.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Okta = Pulumi.Okta;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Okta.App.Bookmark("example", new Okta.App.BookmarkArgs
    ///         {
    ///             Label = "Example",
    ///             Url = "https://example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Bookmark : Pulumi.CustomResource
    {
        /// <summary>
        /// Display auto submit toolbar
        /// </summary>
        [Output("autoSubmitToolbar")]
        public Output<bool?> AutoSubmitToolbar { get; private set; } = null!;

        /// <summary>
        /// Groups associated with the application
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// Do not display application icon on mobile app
        /// </summary>
        [Output("hideIos")]
        public Output<bool?> HideIos { get; private set; } = null!;

        /// <summary>
        /// Do not display application icon to users
        /// </summary>
        [Output("hideWeb")]
        public Output<bool?> HideWeb { get; private set; } = null!;

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// name of app.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Would you like Okta to add an integration for this app?
        /// </summary>
        [Output("requestIntegration")]
        public Output<bool?> RequestIntegration { get; private set; } = null!;

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Output("signOnMode")]
        public Output<string> SignOnMode { get; private set; } = null!;

        /// <summary>
        /// Status of application.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The URL of the bookmark.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.BookmarkUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Bookmark resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bookmark(string name, BookmarkArgs args, CustomResourceOptions? options = null)
            : base("okta:app/bookmark:Bookmark", name, args ?? new BookmarkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bookmark(string name, Input<string> id, BookmarkState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/bookmark:Bookmark", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Bookmark resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bookmark Get(string name, Input<string> id, BookmarkState? state = null, CustomResourceOptions? options = null)
        {
            return new Bookmark(name, id, state, options);
        }
    }

    public sealed class BookmarkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display auto submit toolbar
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Groups associated with the application
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Do not display application icon on mobile app
        /// </summary>
        [Input("hideIos")]
        public Input<bool>? HideIos { get; set; }

        /// <summary>
        /// Do not display application icon to users
        /// </summary>
        [Input("hideWeb")]
        public Input<bool>? HideWeb { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Would you like Okta to add an integration for this app?
        /// </summary>
        [Input("requestIntegration")]
        public Input<bool>? RequestIntegration { get; set; }

        /// <summary>
        /// Status of application.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The URL of the bookmark.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("users")]
        private InputList<Inputs.BookmarkUserArgs>? _users;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        public InputList<Inputs.BookmarkUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.BookmarkUserArgs>());
            set => _users = value;
        }

        public BookmarkArgs()
        {
        }
    }

    public sealed class BookmarkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display auto submit toolbar
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Groups associated with the application
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Do not display application icon on mobile app
        /// </summary>
        [Input("hideIos")]
        public Input<bool>? HideIos { get; set; }

        /// <summary>
        /// Do not display application icon to users
        /// </summary>
        [Input("hideWeb")]
        public Input<bool>? HideWeb { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// name of app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Would you like Okta to add an integration for this app?
        /// </summary>
        [Input("requestIntegration")]
        public Input<bool>? RequestIntegration { get; set; }

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Input("signOnMode")]
        public Input<string>? SignOnMode { get; set; }

        /// <summary>
        /// Status of application.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The URL of the bookmark.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("users")]
        private InputList<Inputs.BookmarkUserGetArgs>? _users;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        public InputList<Inputs.BookmarkUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.BookmarkUserGetArgs>());
            set => _users = value;
        }

        public BookmarkState()
        {
        }
    }
}
