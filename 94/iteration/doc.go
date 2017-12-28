package leetcode

/*
c++ 的迭代实现

vector<int> inorderTraversal(TreeNode* root) {
    vector<int> v;
    if(!root) return v;
    TreeNode* temp = root;
    stack<TreeNode*> s;
    while(true){
        while(temp){
            s.push(temp);
            temp = temp->left;
        }
        if(s.empty()) break;
        temp = s.top();
        s.pop();
        v.push_back(temp->val);
        temp = temp->right;
    }
    return v;
}
*/
