// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Group
{
    /// <summary>
    /// Creates an Okta Group Rule.
    /// 
    /// This resource allows you to create and configure an Okta Group Rule.
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
    ///         var example = new Okta.Group.Rule("example", new Okta.Group.RuleArgs
    ///         {
    ///             ExpressionType = "urn:okta:expression:1.0",
    ///             ExpressionValue = "String.startsWith(user.firstName,\"andy\")",
    ///             GroupAssignments = 
    ///             {
    ///                 "&lt;group id&gt;",
    ///             },
    ///             Status = "ACTIVE",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Rule : Pulumi.CustomResource
    {
        /// <summary>
        /// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
        /// </summary>
        [Output("expressionType")]
        public Output<string?> ExpressionType { get; private set; } = null!;

        /// <summary>
        /// The expression value.
        /// </summary>
        [Output("expressionValue")]
        public Output<string> ExpressionValue { get; private set; } = null!;

        /// <summary>
        /// The list of group ids to assign the users to.
        /// </summary>
        [Output("groupAssignments")]
        public Output<ImmutableArray<string>> GroupAssignments { get; private set; } = null!;

        /// <summary>
        /// The name of the Group Rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the group rule.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("okta:group/rule:Rule", name, args ?? new RuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("okta:group/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
        /// </summary>
        [Input("expressionType")]
        public Input<string>? ExpressionType { get; set; }

        /// <summary>
        /// The expression value.
        /// </summary>
        [Input("expressionValue", required: true)]
        public Input<string> ExpressionValue { get; set; } = null!;

        [Input("groupAssignments", required: true)]
        private InputList<string>? _groupAssignments;

        /// <summary>
        /// The list of group ids to assign the users to.
        /// </summary>
        public InputList<string> GroupAssignments
        {
            get => _groupAssignments ?? (_groupAssignments = new InputList<string>());
            set => _groupAssignments = value;
        }

        /// <summary>
        /// The name of the Group Rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the group rule.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RuleArgs()
        {
        }
    }

    public sealed class RuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
        /// </summary>
        [Input("expressionType")]
        public Input<string>? ExpressionType { get; set; }

        /// <summary>
        /// The expression value.
        /// </summary>
        [Input("expressionValue")]
        public Input<string>? ExpressionValue { get; set; }

        [Input("groupAssignments")]
        private InputList<string>? _groupAssignments;

        /// <summary>
        /// The list of group ids to assign the users to.
        /// </summary>
        public InputList<string> GroupAssignments
        {
            get => _groupAssignments ?? (_groupAssignments = new InputList<string>());
            set => _groupAssignments = value;
        }

        /// <summary>
        /// The name of the Group Rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the group rule.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RuleState()
        {
        }
    }
}
